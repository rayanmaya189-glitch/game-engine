package tournament

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// PrizePool manages tournament prize distribution
type PrizePool struct {
	redisClient *redis.Client
}

// NewPrizePool creates a new prize pool manager
func NewPrizePool(redisClient *redis.Client) *PrizePool {
	return &PrizePool{
		redisClient: redisClient,
	}
}

// CalculatePrizes calculates prize distribution for a tournament
func (p *PrizePool) CalculatePrizes(tournament *Tournament) []Result {
	if len(tournament.RegisteredUsers) == 0 {
		return []Result{}
	}

	// Sort participants by score to determine ranks
	participants := make([]Participant, 0, len(tournament.RegisteredUsers))
	for _, participant := range tournament.RegisteredUsers {
		participants = append(participants, participant)
	}

	// Sort by score descending
	for i := 0; i < len(participants)-1; i++ {
		for j := i + 1; j < len(participants); j++ {
			if participants[j].Score > participants[i].Score {
				participants[i], participants[j] = participants[j], participants[i]
			}
		}
	}

	// Get prize distribution
	prizeDistribution := tournament.Settings.PrizeDistribution
	if len(prizeDistribution) == 0 {
		prizeDistribution = []int{50, 30, 20}
	}

	// Calculate prizes
	results := make([]Result, 0, len(participants))
	prizePool := tournament.PrizePool

	for rank, participant := range participants {
		prize := p.calculatePrize(rank+1, prizeDistribution, prizePool)

		result := Result{
			UserID:    participant.UserID,
			Username:  participant.Username,
			Rank:      rank + 1,
			Prize:     prize,
			Knockouts: participant.Knockouts,
			Score:     participant.Score,
		}
		results = append(results, result)
	}

	return results
}

// calculatePrize calculates the prize for a specific rank
func (p *PrizePool) calculatePrize(rank int, distribution []int, totalPrize int64) int64 {
	if rank > len(distribution) {
		return 0
	}

	percentage := float64(distribution[rank-1]) / 100.0
	return int64(float64(totalPrize) * percentage)
}

// DistributePrizes distributes prizes to winners
func (p *PrizePool) DistributePrizes(ctx context.Context, tournamentID string, results []Result) error {
	key := fmt.Sprintf("prizes:%s", tournamentID)

	for _, result := range results {
		data, err := json.Marshal(result)
		if err != nil {
			return err
		}

		if err := p.redisClient.SAdd(ctx, key, data).Err(); err != nil {
			return err
		}
	}

	return nil
}

// GetPrizeDistribution gets the prize distribution for a tournament
func (p *PrizePool) GetPrizeDistribution(ctx context.Context, tournamentID string) ([]Result, error) {
	key := fmt.Sprintf("prizes:%s", tournamentID)

	members, err := p.redisClient.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	results := make([]Result, 0, len(members))
	for _, member := range members {
		var result Result
		if err := json.Unmarshal([]byte(member), &result); err != nil {
			continue
		}
		results = append(results, result)
	}

	return results, nil
}

// AddPrize adds extra prize money to the pool
func (p *PrizePool) AddPrize(ctx context.Context, tournamentID string, amount int64) error {
	key := fmt.Sprintf("prizepool:%s", tournamentID)

	return p.redisClient.IncrBy(ctx, key, amount).Err()
}

// GetTotalPrize gets the total prize pool for a tournament
func (p *PrizePool) GetTotalPrize(ctx context.Context, tournamentID string) (int64, error) {
	key := fmt.Sprintf("prizepool:%s", tournamentID)

	prize, err := p.redisClient.Get(ctx, key).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return prize, err
}

// InitializePrizePool initializes the prize pool for a tournament
func (p *PrizePool) InitializePrizePool(ctx context.Context, tournamentID string, amount int64) error {
	key := fmt.Sprintf("prizepool:%s", tournamentID)

	return p.redisClient.Set(ctx, key, amount, 0).Err()
}

// GetPrizeBreakdown returns detailed prize breakdown
func (p *PrizePool) GetPrizeBreakdown(distribution []int, totalPrize int64) []struct {
	Rank    int
	Min     int
	Max     int
	Amount  int64
	Percent int
} {
	breakdown := make([]struct {
		Rank    int
		Min     int
		Max     int
		Amount  int64
		Percent int
	}, 0, len(distribution))

	prevMax := 0
	for i, percent := range distribution {
		rank := i + 1
		min := prevMax + 1
		max := rank * len(distribution) / len(distribution)
		if max < min {
			max = min
		}
		prevMax = max

		amount := int64(float64(totalPrize) * float64(percent) / 100.0)

		breakdown = append(breakdown, struct {
			Rank    int
			Min     int
			Max     int
			Amount  int64
			Percent int
		}{
			Rank:    rank,
			Min:     min,
			Max:     max,
			Amount:  amount,
			Percent: percent,
		})
	}

	return breakdown
}

// CalculateBounty calculates bounty for knockout tournaments
func (p *PrizePool) CalculateBounty(buyIn int64) int64 {
	// Bounty is typically a percentage of the buy-in
	return buyIn / 2
}

// AddBounty adds bounty prize for a knockout
func (p *PrizePool) AddBounty(ctx context.Context, tournamentID string, userID string, amount int64) error {
	key := fmt.Sprintf("bounties:%s", tournamentID)

	return p.redisClient.HIncrBy(ctx, key, userID, amount).Err()
}

// GetBounties gets all bounties for a tournament
func (p *PrizePool) GetBounties(ctx context.Context, tournamentID string) (map[string]int64, error) {
	key := fmt.Sprintf("bounties:%s", tournamentID)

	return p.redisClient.HGetAll(ctx, key).Result()
}

// GetPlayerBounty gets a specific player's bounty
func (p *PrizePool) GetPlayerBounty(ctx context.Context, tournamentID string, userID string) (int64, error) {
	key := fmt.Sprintf("bounties:%s", tournamentID)

	bounty, err := p.redisClient.HGet(ctx, key, userID).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return bounty, err
}

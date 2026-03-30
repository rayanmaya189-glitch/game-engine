package room

import (
	"context"
	"time"
)

// FindMatches finds matching players for a given player
func (m *Matchmaker) FindMatches(ctx context.Context, userID, username, gameType string, skillRating, stakeLevel int) []string {
	key := m.getQueueKey(gameType, stakeLevel)

	m.mu.RLock()
	queue, ok := m.queues[key]
	m.mu.RUnlock()

	if !ok {
		return nil
	}

	player := &MatchmakingPlayer{
		UserID:      userID,
		Username:    username,
		SkillRating: skillRating,
		StakeLevel:  stakeLevel,
		GameType:    gameType,
		WaitStart:   time.Now(),
	}

	matches := queue.FindMatch(ctx, player)

	result := make([]string, len(matches))
	for i, m := range matches {
		result[i] = m.UserID
	}

	return result
}

// AutoMatch attempts to automatically form a game
func (m *Matchmaker) AutoMatch(ctx context.Context, gameType string, stakeLevel int) []string {
	key := m.getQueueKey(gameType, stakeLevel)

	m.mu.RLock()
	queue, ok := m.queues[key]
	m.mu.RUnlock()

	if !ok {
		return nil
	}

	queue.mu.RLock()
	var players []*MatchmakingPlayer
	for _, p := range queue.players {
		players = append(players, p)
	}
	queue.mu.RUnlock()

	if len(players) < m.config.MinPlayers {
		return nil
	}

	var matched []string
	skillSum := 0

	for _, p := range players {
		if len(matched) >= m.config.MaxPlayers {
			break
		}

		if len(matched) == 0 {
			matched = append(matched, p.UserID)
			skillSum = p.SkillRating
		} else {
			avgSkill := skillSum / len(matched)
			if abs(p.SkillRating-avgSkill) <= m.config.MaxSkillDiff {
				matched = append(matched, p.UserID)
				skillSum += p.SkillRating
			}
		}
	}

	if len(matched) < m.config.MinPlayers {
		firstPlayer := players[0]
		if time.Since(firstPlayer.WaitStart) > m.config.MaxWaitTime && len(matched) >= 2 {
			return matched
		}
		return nil
	}

	for _, userID := range matched {
		queue.RemovePlayer(ctx, userID)
	}

	return matched
}

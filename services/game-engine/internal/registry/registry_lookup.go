package registry

// ListByType returns game definitions filtered by type
func (r *GameRegistry) ListByType(gameType string) []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*GameDefinition
	for _, game := range r.games {
		if game.Type == gameType {
			result = append(result, game)
		}
	}

	return result
}

// GetActiveGames returns all active games
func (r *GameRegistry) GetActiveGames() []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*GameDefinition
	for _, game := range r.games {
		if game.Status == "active" {
			result = append(result, game)
		}
	}

	return result
}

// DefaultGames returns default game definitions
func DefaultGames() []*GameDefinition {
	return []*GameDefinition{
		{
			ID:          "blackjack",
			Name:        "Blackjack",
			Type:        "card",
			Subtype:     "blackjack",
			Description: "Classic blackjack game",
			RTP:         99.5,
			HouseEdge:   0.5,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "baccarat",
			Name:        "Baccarat",
			Type:        "card",
			Subtype:     "baccarat",
			Description: "Classic baccarat game",
			RTP:         98.94,
			HouseEdge:   1.06,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "roulette",
			Name:        "European Roulette",
			Type:        "table",
			Subtype:     "roulette",
			Description: "European roulette with single zero",
			RTP:         97.3,
			HouseEdge:   2.7,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   35000,
			Status:      "active",
		},
		{
			ID:          "slots_classic",
			Name:        "Classic Slots",
			Type:        "slot",
			Subtype:     "classic",
			Description: "3-reel classic slot machine",
			RTP:         95.5,
			HouseEdge:   4.5,
			MinBet:      1,
			MaxBet:      100,
			MaxPayout:   1000,
			Status:      "active",
		},
		{
			ID:          "slots_video",
			Name:        "Video Slots",
			Type:        "slot",
			Subtype:     "video",
			Description: "5-reel video slot machine",
			RTP:         96.5,
			HouseEdge:   3.5,
			MinBet:      1,
			MaxBet:      500,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "craps",
			Name:        "Craps",
			Type:        "dice",
			Subtype:     "craps",
			Description: "Classic craps table game",
			RTP:         98.5,
			HouseEdge:   1.5,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "sicbo",
			Name:        "Sic Bo",
			Type:        "dice",
			Subtype:     "sicbo",
			Description: "Chinese dice game",
			RTP:         97.5,
			HouseEdge:   2.5,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "poker",
			Name:        "Texas Hold'em",
			Type:        "card",
			Subtype:     "poker",
			Description: "Texas Hold'em poker",
			RTP:         96.0,
			HouseEdge:   4.0,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   100000,
			Status:      "active",
		},
		{
			ID:          "dragon_tiger",
			Name:        "Dragon Tiger",
			Type:        "card",
			Subtype:     "dragon_tiger",
			Description: "Dragon Tiger card game",
			RTP:         97.5,
			HouseEdge:   2.5,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "casino_war",
			Name:        "Casino War",
			Type:        "card",
			Subtype:     "casino_war",
			Description: "Casino War card game",
			RTP:         97.0,
			HouseEdge:   3.0,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "three_card_poker",
			Name:        "Three Card Poker",
			Type:        "card",
			Subtype:     "three_card_poker",
			Description: "Three Card Poker game",
			RTP:         96.0,
			HouseEdge:   4.0,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "andar_bahar",
			Name:        "Andar Bahar",
			Type:        "card",
			Subtype:     "andar_bahar",
			Description: "Andar Bahar card game",
			RTP:         97.0,
			HouseEdge:   3.0,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "teen_patti",
			Name:        "Teen Patti",
			Type:        "card",
			Subtype:     "teen_patti",
			Description: "Teen Patti card game",
			RTP:         96.5,
			HouseEdge:   3.5,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "hilo",
			Name:        "Hi-Lo",
			Type:        "dice",
			Subtype:     "hilo",
			Description: "High-Low dice game",
			RTP:         97.0,
			HouseEdge:   3.0,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "slots_megaways",
			Name:        "Megaways Slots",
			Type:        "slot",
			Subtype:     "megaways",
			Description: "Megaways slot machine with up to 117,649 ways to win",
			RTP:         96.0,
			HouseEdge:   4.0,
			MinBet:      1,
			MaxBet:      500,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "slots_cluster",
			Name:        "Cluster Pays Slots",
			Type:        "slot",
			Subtype:     "cluster",
			Description: "Cluster Pays slot machine",
			RTP:         96.0,
			HouseEdge:   4.0,
			MinBet:      1,
			MaxBet:      500,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "slots_progressive",
			Name:        "Progressive Jackpot Slots",
			Type:        "slot",
			Subtype:     "progressive",
			Description: "Progressive jackpot slot machine",
			RTP:         94.0,
			HouseEdge:   6.0,
			MinBet:      1,
			MaxBet:      100,
			MaxPayout:   1000000,
			Status:      "active",
		},
	}
}

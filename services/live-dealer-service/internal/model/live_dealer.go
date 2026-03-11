package model

// Table represents a live dealer table
type Table struct {
	TableID     string `json:"table_id"`
	GameType    string `json:"game_type"`
	DealerName  string `json:"dealer_name"`
	Status      string `json:"status"` // open, closed, maintenance
	MinBet      int    `json:"min_bet"`
	MaxBet      int    `json:"max_bet"`
	CurrentSeat int    `json:"current_seat"`
	MaxSeats    int    `json:"max_seats"`
}

// Player represents a player at a table
type Player struct {
	PlayerID   string  `json:"player_id"`
	TableID    string  `json:"table_id"`
	SeatNumber int     `json:"seat_number"`
	Chips      float64 `json:"chips"`
	JoinedAt   int64   `json:"joined_at"`
}

// GameState represents the current state of a game
type GameState struct {
	TableID    string   `json:"table_id"`
	RoundID    string   `json:"round_id"`
	Phase      string   `json:"phase"` // betting, playing, resolving
	Cards      []string `json:"cards"`
	DealerCard string   `json:"dealer_card"`
	Pot        float64  `json:"pot"`
	UpdatedAt  int64    `json:"updated_at"`
}

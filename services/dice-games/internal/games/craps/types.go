package craps

import "time"

// Game phases
const (
	PhaseComeOut = "come_out"
	PhasePoint   = "point"
	PhaseRoll    = "roll"
)

// Bet types
const (
	BetPassLine    = "pass_line"
	BetDontPass    = "dont_pass"
	BetCome        = "come"
	BetDontCome    = "dont_come"
	BetPlaceWin4   = "place_win_4"
	BetPlaceWin5   = "place_win_5"
	BetPlaceWin6   = "place_win_6"
	BetPlaceWin8   = "place_win_8"
	BetPlaceWin9   = "place_win_9"
	BetPlaceWin10  = "place_win_10"
	BetPlaceLose4  = "place_lose_4"
	BetPlaceLose5  = "place_lose_5"
	BetPlaceLose6  = "place_lose_6"
	BetPlaceLose8  = "place_lose_8"
	BetPlaceLose9  = "place_lose_9"
	BetPlaceLose10 = "place_lose_10"
	BetField       = "field"
	BetBig6        = "big_6"
	BetBig8        = "big_8"
	BetAny7        = "any_7"
	BetAnyCraps    = "any_craps"
	BetHorn2       = "horn_2"
	BetHorn3       = "horn_3"
	BetHorn11      = "horn_11"
	BetHorn12      = "horn_12"
	BetHard4       = "hard_4"
	BetHard6       = "hard_6"
	BetHard8       = "hard_8"
	BetHard10      = "hard_10"
	BetBuy4        = "buy_4"
	BetBuy5        = "buy_5"
	BetBuy6        = "buy_6"
	BetBuy8        = "buy_8"
	BetBuy9        = "buy_9"
	BetBuy10       = "buy_10"
	BetLay4        = "lay_4"
	BetLay5        = "lay_5"
	BetLay6        = "lay_6"
	BetLay8        = "lay_8"
	BetLay9        = "lay_9"
	BetLay10       = "lay_10"
)

// Payout ratios for each bet type
var payoutRatios = map[string]float64{
	BetPassLine:    1.0,
	BetDontPass:    1.0,
	BetCome:        1.0,
	BetDontCome:    1.0,
	BetPlaceWin4:   9.0 / 5.0,  // 9:5
	BetPlaceWin5:   7.0 / 5.0,  // 7:5
	BetPlaceWin6:   7.0 / 6.0,  // 7:6
	BetPlaceWin8:   7.0 / 6.0,  // 7:6
	BetPlaceWin9:   7.0 / 5.0,  // 7:5
	BetPlaceWin10:  9.0 / 5.0,  // 9:5
	BetPlaceLose4:  5.0 / 11.0, // 5:11
	BetPlaceLose5:  5.0 / 8.0,  // 5:8
	BetPlaceLose6:  4.0 / 5.0,  // 4:5
	BetPlaceLose8:  4.0 / 5.0,  // 4:5
	BetPlaceLose9:  5.0 / 8.0,  // 5:8
	BetPlaceLose10: 5.0 / 11.0, // 5:11
	BetField:       1.0,        // 1:1 (2 and 12 pay 2:1)
	BetBig6:        1.0,
	BetBig8:        1.0,
	BetAny7:        4.0,
	BetAnyCraps:    7.0,
	BetHorn2:       27.0 / 4.0, // 27:4
	BetHorn3:       3.0,
	BetHorn11:      3.0,
	BetHorn12:      27.0 / 4.0, // 27:4
	BetHard4:       7.0,
	BetHard6:       7.0 / 6.0, // 7:6
	BetHard8:       7.0 / 6.0, // 7:6
	BetHard10:      7.0,
	BetBuy4:        2.0, // 5% commission
	BetBuy5:        2.0,
	BetBuy6:        2.0,
	BetBuy8:        2.0,
	BetBuy9:        2.0,
	BetBuy10:       2.0,
	BetLay4:        1.0 / 2.0, // 1:2
	BetLay5:        2.0 / 3.0, // 2:3
	BetLay6:        5.0 / 6.0, // 5:6
	BetLay8:        5.0 / 6.0, // 5:6
	BetLay9:        2.0 / 3.0, // 2:3
	BetLay10:       1.0 / 2.0, // 1:2
}

// CrapsNumbers identifies craps numbers (2, 3, 12)
var crapsNumbers = map[int]bool{2: true, 3: true, 12: true}

// PointNumbers identifies point numbers
var pointNumbers = map[int]bool{4: true, 5: true, 6: true, 8: true, 9: true, 10: true}

// Game represents a Craps game
type Game struct {
	ID           string
	Point        int
	Dice         []int
	Phase        string
	Bets         map[string]map[string]int64 // betType -> playerID -> amount
	Payouts      map[string]map[string]int64 // betType -> playerID -> payout
	Roller       Roller
	ProvablyFair bool
	ServerSeed   string
	ClientSeed   string
	Nonce        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Roller interface for dice rolling
type Roller interface {
	Roll() ([]int, error)
}

// GameState represents the game state
type GameState struct {
	GameID         string                      `json:"game_id"`
	Dice           []int                       `json:"dice"`
	Sum            int                         `json:"sum"`
	Point          int                         `json:"point"`
	Phase          string                      `json:"phase"`
	Bets           map[string]map[string]int64 `json:"bets"`
	Payouts        map[string]map[string]int64 `json:"payouts"`
	ProvablyFair   bool                        `json:"provably_fair"`
	ServerSeedHash string                      `json:"server_seed_hash,omitempty"`
	ClientSeed     string                      `json:"client_seed,omitempty"`
	Nonce          int                         `json:"nonce,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

// Bet information maps
var betNames = map[string]string{
	BetPassLine:    "Pass Line",
	BetDontPass:    "Don't Pass",
	BetCome:        "Come",
	BetDontCome:    "Don't Come",
	BetPlaceWin4:   "Place 4 Win",
	BetPlaceWin5:   "Place 5 Win",
	BetPlaceWin6:   "Place 6 Win",
	BetPlaceWin8:   "Place 8 Win",
	BetPlaceWin9:   "Place 9 Win",
	BetPlaceWin10:  "Place 10 Win",
	BetPlaceLose4:  "Place 4 Lose",
	BetPlaceLose5:  "Place 5 Lose",
	BetPlaceLose6:  "Place 6 Lose",
	BetPlaceLose8:  "Place 8 Lose",
	BetPlaceLose9:  "Place 9 Lose",
	BetPlaceLose10: "Place 10 Lose",
	BetField:       "Field",
	BetBig6:        "Big 6",
	BetBig8:        "Big 8",
	BetAny7:        "Any 7",
	BetAnyCraps:    "Any Craps",
	BetHorn2:       "Horn 2",
	BetHorn3:       "Horn 3",
	BetHorn11:      "Horn 11",
	BetHorn12:      "Horn 12",
	BetHard4:       "Hard 4",
	BetHard6:       "Hard 6",
	BetHard8:       "Hard 8",
	BetHard10:      "Hard 10",
	BetBuy4:        "Buy 4",
	BetBuy5:        "Buy 5",
	BetBuy6:        "Buy 6",
	BetBuy8:        "Buy 8",
	BetBuy9:        "Buy 9",
	BetBuy10:       "Buy 10",
	BetLay4:        "Lay 4",
	BetLay5:        "Lay 5",
	BetLay6:        "Lay 6",
	BetLay8:        "Lay 8",
	BetLay9:        "Lay 9",
	BetLay10:       "Lay 10",
}

var betDescriptions = map[string]string{
	BetPassLine: "Win on 7 or 11 (come-out) or match point",
	BetDontPass: "Win on 2, 3, or 7 before point",
	BetField:    "Win on 3,4,9,10,11 (2 and 12 pay double)",
	BetBig6:     "Win when 6 is rolled before 7",
	BetBig8:     "Win when 8 is rolled before 7",
	BetAny7:     "Win when any 7 is rolled",
	BetAnyCraps: "Win when 2, 3, or 12 is rolled",
	BetHard4:    "Win on double 2 (hard 4) before 7",
	BetHard6:    "Win on double 3 (hard 6) before 7",
	BetHard8:    "Win on double 4 (hard 8) before 7",
	BetHard10:   "Win on double 5 (hard 10) before 7",
}

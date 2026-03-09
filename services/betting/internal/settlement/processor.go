package settlement

// Result represents a bet result
type Result int

const (
	Win Result = iota
	Loss
	Push
	Cancelled
)

// Processor handles bet settlement
type Processor struct{}

// NewProcessor creates a new settlement processor
func NewProcessor() *Processor {
	return &Processor{}
}

// Settle settles a bet
func (p *Processor) Settle(result Result, stake int64, odds float64) int64 {
	switch result {
	case Win:
		return int64(float64(stake) * odds)
	case Push:
		return stake
	case Loss, Cancelled:
		return 0
	}
	return 0
}

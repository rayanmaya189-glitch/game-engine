package blackjack

// settle settles all bets
func (g *Game) settle() {
	dealerTotal := calculateTotal(g.Dealer.Hand)
	dealerBusted := dealerTotal > 21

	for _, player := range g.Players {
		for _, hand := range player.Hands {
			if hand.Result != ResultPending && hand.Result != ResultBlackjack {
				continue
			}

			playerTotal := calculateTotal(hand.Cards)

			if dealerBusted {
				if playerTotal <= 21 {
					hand.Result = ResultWin
				} else {
					hand.Result = ResultLoss
				}
			} else {
				if playerTotal > 21 {
					hand.Result = ResultLoss
				} else if playerTotal > dealerTotal {
					hand.Result = ResultWin
				} else if playerTotal < dealerTotal {
					hand.Result = ResultLoss
				} else {
					hand.Result = ResultPush
				}
			}
		}
	}
}

// settleDealerBlackjack settles the game when dealer has blackjack
func (g *Game) settleDealerBlackjack() {
	for _, player := range g.Players {
		for _, hand := range player.Hands {
			if hand.Result == ResultBlackjack {
				hand.Result = ResultPush
			} else if hand.Result == ResultPending {
				hand.Result = ResultLoss
			}
		}
	}
}

// GetWinnings calculates the winnings for a hand
func GetWinnings(hand *Hand, payout float64) int64 {
	if hand.Result != ResultWin && hand.Result != ResultBlackjack {
		return -hand.Bet
	}

	if hand.Result == ResultBlackjack {
		return int64(float64(hand.Bet) * (1 + payout))
	}

	return hand.Bet
}

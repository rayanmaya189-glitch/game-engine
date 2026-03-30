package games

func (g *Game) checkScatters() {
	scatterCount := 0
	for reel := 0; reel < g.Config.Reels; reel++ {
		for row := 0; row < g.Config.Rows; row++ {
			if g.Symbols[g.Positions[reel][row]].IsScatter {
				scatterCount++
			}
		}
	}

	if scatterCount >= 3 {
		g.ScatterWin = int64(scatterCount) * 10 * g.LineBet
	}
}

func (g *Game) checkBonus() {
	bonusCount := 0
	for reel := 0; reel < g.Config.Reels; reel++ {
		for row := 0; row < g.Config.Rows; row++ {
			if g.Symbols[g.Positions[reel][row]].IsBonus {
				bonusCount++
			}
		}
	}

	if bonusCount >= 3 {
		g.FreeSpins = 10
		g.FreeSpinsLeft = 10
	}
}

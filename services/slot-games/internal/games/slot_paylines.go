package games

func createPaylines(reels, rows int) []Payline {
	paylines := []Payline{}

	// Horizontal paylines
	for r := 0; r < rows; r++ {
		positions := make([][]int, reels)
		for i := 0; i < reels; i++ {
			positions[i] = []int{r}
		}
		paylines = append(paylines, Payline{
			ID:        len(paylines) + 1,
			Name:      "Row " + string(r+'1'),
			Positions: positions,
			Payout:    100,
			MinSymbol: 3,
		})
	}

	// Zigzag paylines
	if rows == 3 && reels >= 3 {
		positions1 := [][]int{{0, 0}, {1, 1}, {2, 2}}
		paylines = append(paylines, Payline{ID: len(paylines) + 1, Name: "Diagonal 1", Positions: positions1, Payout: 200, MinSymbol: 3})

		positions2 := [][]int{{0, 2}, {1, 1}, {2, 0}}
		paylines = append(paylines, Payline{ID: len(paylines) + 1, Name: "Diagonal 2", Positions: positions2, Payout: 200, MinSymbol: 3})
	}

	return paylines
}

// evaluate evaluates the spin result
func (g *Game) evaluate() {
	g.Win = 0
	g.WinLines = nil
	g.ScatterWin = 0
	g.BonusWin = 0
	g.TotalWin = 0
	g.FreeSpins = 0

	for i := 0; i < g.Lines && i < len(g.Paylines); i++ {
		winLine := g.checkPayline(g.Paylines[i])
		if winLine != nil {
			g.WinLines = append(g.WinLines, *winLine)
			g.Win += winLine.Payout
		}
	}

	g.checkScatters()
	g.checkBonus()

	g.TotalWin = (g.Win + g.ScatterWin + g.BonusWin) * g.LineBet
}

func (g *Game) checkPayline(payline Payline) *WinLine {
	symbols := make([]int, len(payline.Positions))
	for i, pos := range payline.Positions {
		if pos[0] < len(g.Positions) && pos[1] < len(g.Positions[pos[0]]) {
			symbols[i] = g.Positions[pos[0]][pos[1]]
		} else {
			return nil
		}
	}

	matchCount := 1
	matchingSymbol := symbols[0]

	for i := 1; i < len(symbols); i++ {
		isWild := g.Symbols[symbols[i]].IsWild || g.Symbols[matchingSymbol].IsWild
		if symbols[i] == matchingSymbol || isWild {
			matchCount++
		} else {
			break
		}
	}

	if matchCount >= payline.MinSymbol {
		payout := int64(payline.Payout)
		for i := 3; i < matchCount; i++ {
			payout = payout * 2
		}

		return &WinLine{
			PaylineID: payline.ID,
			Symbol:    g.Symbols[matchingSymbol].ID,
			Count:     matchCount,
			Payout:    payout,
		}
	}

	return nil
}

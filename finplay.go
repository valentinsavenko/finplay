package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot/plotter"
)

// a scenario that generates 4 plotable slices of tupels
type goldScenario func() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs)
type goldPriceModifier func(index int, goldPrice float64) float64

const (
	goldPrice               = 1500.0 // dollar
	colletralizedPercentage = 80.0
	rebuys                  = 100 // times you'll rebuy another gold coin
)

func goldPlay(goldMod goldPriceModifier) (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {

	goldValues := make(plotter.XYs, rebuys)
	ownMoneyInvested := make(plotter.XYs, rebuys)
	loanedMoney := make(plotter.XYs, rebuys)
	priceOfAssets := make(plotter.XYs, rebuys)

	currentGoldPrice := goldPrice

	for i := range goldValues {
		currentGoldPrice = goldMod(i, currentGoldPrice)

		goldValues[i].X = float64(i)
		goldValues[i].Y = currentGoldPrice

		ownMoneyInvested[i].X = float64(i)
		additionalOwnMoney := (currentGoldPrice / 100) * (100 - colletralizedPercentage)
		if i == 0 {
			ownMoneyInvested[i].Y = currentGoldPrice
		} else {
			ownMoneyInvested[i].Y = ownMoneyInvested[i-1].Y + additionalOwnMoney
		}

		loanedMoney[i].X = float64(i)
		currentMoneyLoan := (currentGoldPrice / 100) * colletralizedPercentage
		if i == 0 {
			loanedMoney[i].Y = 0
		} else {
			loanedMoney[i].Y = loanedMoney[i-1].Y + currentMoneyLoan
		}

		priceOfAssets[i].X = float64(i)
		priceOfAssets[i].Y = currentGoldPrice * float64(i)

	}
	return goldValues, ownMoneyInvested, loanedMoney, priceOfAssets
}

func constantPrice(index int, goldPrice float64) float64 {
	return goldPrice
}

func naiveGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(constantPrice)
}

func fallingPrice(index int, goldPrice float64) float64 {
	if index < 20 {
		return goldPrice
	}
	return (goldPrice / 100) * 95
}

func unluckyGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(fallingPrice)
}

func raisingPrice(index int, goldPrice float64) float64 {
	if index < 20 {
		return goldPrice
	}
	return (goldPrice / 100) * 105
}

func luckyGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(raisingPrice)
}

func fluctuatingPrice(index int, goldPrice float64) float64 {
	// based on index, this number functuates between -10 and +10
	fluctuation10 := 10 * math.Sin(float64(index))
	return (goldPrice / 100) * (100 + fluctuation10)
}

func fluctuatingGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(fluctuatingPrice)
}

func main() {

	plotGoldStrategy(naiveGoldPlay, "testdata/naive_rebuy.png")
	plotGoldStrategy(unluckyGoldPlay, "testdata/unlucky_rebuy.png")
	plotGoldStrategy(luckyGoldPlay, "testdata/lucky_rebuy.png")
	plotGoldStrategy(fluctuatingGoldPlay, "testdata/fluctuating_rebuy.png")

	// historical gold Data
	goldPriceData := parseCSV("historical_gold_price.csv")
	pTimed := getTimeSeriesPlotter()
	green := color.RGBA{R: 2, G: 166, B: 67, A: 255}

	goldTimedLine := getColoredLine(goldPriceData, green)
	pTimed.Add(goldTimedLine)

	savePlotTo(pTimed, "testdata/timeseries.png")
}

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
	previousGoldPrice := goldPrice

	for i := range goldValues {
		currentGoldPrice = goldMod(i, currentGoldPrice)

		goldValues[i].X = float64(i)
		goldValues[i].Y = currentGoldPrice

		loanedMoney[i].X = float64(i)
		var currentMoneyLoan float64 = 0
		if i != 0 {
			currentMoneyLoan = (previousGoldPrice / 100) * colletralizedPercentage
			loanedMoney[i].Y = loanedMoney[i-1].Y + currentMoneyLoan
		}

		ownMoneyInvested[i].X = float64(i)
		additionalOwnMoney := currentGoldPrice - currentMoneyLoan
		if i == 0 {
			ownMoneyInvested[i].Y = additionalOwnMoney
		} else {
			ownMoneyInvested[i].Y = ownMoneyInvested[i-1].Y + additionalOwnMoney
		}

		priceOfAssets[i].X = float64(i)
		priceOfAssets[i].Y = currentGoldPrice * float64(i)

		previousGoldPrice = currentGoldPrice
	}
	return goldValues, ownMoneyInvested, loanedMoney, priceOfAssets
}

// -----------------------------------------
// naive gold scenario
// -----------------------------------------
func constantPrice(index int, goldPrice float64) float64 {
	return goldPrice
}

func naiveGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(constantPrice)
}

// -----------------------------------------
// falling gold price scenario
// -----------------------------------------
func fallingPrice(index int, goldPrice float64) float64 {
	if index < 20 {
		return goldPrice
	}
	return (goldPrice / 100) * 95
}

func unluckyGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(fallingPrice)
}

// -----------------------------------------
// increasing gold price scenario
// -----------------------------------------
func increasingPrice(index int, goldPrice float64) float64 {
	if index < 20 {
		return goldPrice
	}
	return (goldPrice / 100) * 105
}

func luckyGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(increasingPrice)
}

// -----------------------------------------
// swinging (realistic) gold price scenario
// -----------------------------------------
func fluctuatingPrice(index int, goldPrice float64) float64 {
	// based on index, this number functuates between -10 and +10
	fluctuation10 := 10 * math.Sin(float64(index))
	return (goldPrice / 100) * (100 + fluctuation10)
}

func fluctuatingGoldPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {
	return goldPlay(fluctuatingPrice)
}

func main() {

	plotGoldStrategy(naiveGoldPlay, "plotted_graphs/naive_rebuy.png")
	plotGoldStrategy(unluckyGoldPlay, "plotted_graphs/unlucky_rebuy.png")
	plotGoldStrategy(luckyGoldPlay, "plotted_graphs/lucky_rebuy.png")
	plotGoldStrategy(fluctuatingGoldPlay, "plotted_graphs/fluctuating_rebuy.png")

	// historical gold Data
	goldPriceData := parseCSV("historical_gold_price.csv")
	pTimed := getTimeSeriesPlotter()
	green := color.RGBA{R: 2, G: 166, B: 67, A: 255}

	goldTimedLine := getColoredLine(goldPriceData, green)
	pTimed.Add(goldTimedLine)

	savePlotTo(pTimed, "plotted_graphs/timeseries.png")
}

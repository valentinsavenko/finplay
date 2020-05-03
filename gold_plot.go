package main

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func getRebuyPlotter() *plot.Plot {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "FINPLAY"
	p.X.Label.Text = "#coins"
	p.Y.Label.Text = "USD"
	p.Add(plotter.NewGrid())

	return p
}

func getTimeSeriesPlotter() *plot.Plot {
	xticks := plot.TimeTicks{Format: "2006-01-02\n15:04"}

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "FINPLAY"

	p.X.Tick.Marker = xticks
	p.X.Label.Text = "date"

	p.Y.Label.Text = "USD"
	p.Add(plotter.NewGrid())

	return p
}

func savePlotTo(p *plot.Plot, filename string) {
	err := p.Save(20*vg.Centimeter, 12*vg.Centimeter, filename)
	if err != nil {
		log.Panic(err)
	}
}

func getColoredLine(data plotter.XYs, c color.Color) *plotter.Line {
	line, err := plotter.NewLine(data)
	if err != nil {
		log.Panic(err)
	}
	line.Color = c

	return line
}

func plotGoldStrategy(playGoldScenario goldScenario, filename string) {
	// pre-set colors
	gold := color.RGBA{R: 217, G: 177, B: 0, A: 255}
	green := color.RGBA{R: 2, G: 166, B: 67, A: 255}
	red := color.RGBA{R: 224, G: 70, B: 4, A: 255}
	blue := color.RGBA{R: 27, G: 0, B: 179, A: 255}

	goldValues, ownMoneyInvested, loanedMoney, priceOfAssets := playGoldScenario()
	p := getRebuyPlotter()
	lineGold := getColoredLine(goldValues, gold)
	lineGreen := getColoredLine(ownMoneyInvested, green)
	lineRed := getColoredLine(loanedMoney, red)
	lineBlue := getColoredLine(priceOfAssets, blue)

	p.Add(lineGold, lineGreen, lineRed, lineBlue)
	p.Legend.Add("gold price", lineGold)
	p.Legend.Add("own money invested", lineGreen)
	p.Legend.Add("loaned money", lineRed)
	p.Legend.Add("price of all assets", lineBlue)
	p.Legend.Top = true
	p.Legend.Left = true

	savePlotTo(p, filename)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"gonum.org/v1/plot/plotter"
)

func testPlay() (plotter.XYs, plotter.XYs, plotter.XYs, plotter.XYs) {

	n := 2
	goldValues := make(plotter.XYs, n)
	ownMoneyInvested := make(plotter.XYs, n)
	loanedMoney := make(plotter.XYs, n)
	priceOfAssets := make(plotter.XYs, n)

	goldValues[0].X = 1.0
	goldValues[0].Y = 1.0
	goldValues[1].X = 2.0
	goldValues[1].Y = 30.0

	ownMoneyInvested[0].X = 1.0
	ownMoneyInvested[0].Y = 2.0
	ownMoneyInvested[1].X = 2.0
	ownMoneyInvested[1].Y = 40.0

	loanedMoney[0].X = 1.0
	loanedMoney[0].Y = 3.0
	loanedMoney[1].X = 2.0
	loanedMoney[1].Y = 50.0

	priceOfAssets[0].X = 1.0
	priceOfAssets[0].Y = 4.0
	priceOfAssets[1].X = 2.0
	priceOfAssets[1].Y = 60.0

	return goldValues, ownMoneyInvested, loanedMoney, priceOfAssets
}

func TestPlotGoldStrategy(t *testing.T) {

	resultPNG := "testdata/test.png"
	plotGoldStrategy(testPlay, resultPNG)

	// Open File
	f, err := os.Open(resultPNG)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Get the content
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content Type: " + contentType)

	expectedFileType := "image/png"

	if contentType != expectedFileType {
		msg := "Expected the test.png to be a valid image of type %f but it was %f"
		t.Fatalf(msg, expectedFileType, contentType)
	}
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

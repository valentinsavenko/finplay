package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot/plotter"
)

const (
	layoutISO = "1/2/2006"
	layoutDE  = "2. January 2006"
)

func parseCSV(filename string) plotter.XYs {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Read() // ignore first line

	var goldPrices plotter.XYs
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		parsedDate, _ := time.Parse(layoutISO, line[0])

		var parsedPrice float64
		stringPrice := line[1]
		if s, err := strconv.ParseFloat(stringPrice, 64); err == nil {
			parsedPrice = s
		} else {
			log.Fatal(err)
		}

		var newPoint plotter.XY
		newPoint.X = float64(parsedDate.Unix())
		newPoint.Y = parsedPrice
		goldPrices = append(goldPrices, newPoint)
	}
	return goldPrices
}

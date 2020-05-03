package main

import (
	"testing"
)

func TestParseCSV(t *testing.T) {

	data := parseCSV("testdata/test.csv")
	actual := data[2]
	expectedX := 284083200.000000 // unit time for 1/2/1979
	expectedY := 226.8

	if actual.X != expectedX {
		msg := "Expected third row of test csv be parsed to %f but it was %f"
		t.Fatalf(msg, expectedX, actual.X)

	}
	if actual.Y != expectedY {
		msg := "Expected third row of test csv be parsed to %f but it was %f"
		t.Fatalf(msg, expectedY, actual.Y)
	}
}

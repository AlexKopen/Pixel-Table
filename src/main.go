package main

import "fmt"

func main() {
	var streamEmissions []StreamEmission

	if UseTestData {
		processStreamData(TestStreams, "UNI")
	} else {
		//	Read stream data for each symbol
		for _, symbol := range Symbols {
			fmt.Println(symbol)
			// Process newly acquired stream emissions
			processStreamData(streamEmissions, symbol)
		}
	}
}

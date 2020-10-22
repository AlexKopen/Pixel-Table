package main

func main() {
	if UseTestData {
		processStreamData(TestStreams, "UNI")
	} else {
		//	Read stream data for each symbol
		for _, symbol := range Symbols {
			// Process newly acquired stream emissions
			processStreamData(fetchCoinData(symbol), symbol)
		}
	}
}

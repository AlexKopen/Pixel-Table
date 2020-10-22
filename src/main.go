package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Wait group for concurrency
	if UseTestData {
		//processStreamData(TestStreams, "UNI")
	} else {
		//	Read stream data for each symbol
		for _, symbol := range Symbols {
			// Process newly acquired stream emissions
			streamGenerationChannel := make(chan []StreamEmission)
			wg.Add(1)
			go generateStreamEmissions(symbol, streamGenerationChannel)
			go receiveStreamGenerationOutput(streamGenerationChannel, symbol)
		}

		wg.Wait()
	}
}

func receiveStreamGenerationOutput(c chan []StreamEmission, symbol string) {
	streams := <-c
	processStreamChannel := make(chan float64)
	log.Printf("Streams fetched: %s\n", symbol)
	go processStreamData(streams, symbol, processStreamChannel)
	go receiveProcessStreamDataOutput(symbol, processStreamChannel)
}

func receiveProcessStreamDataOutput(symbol string, c chan float64) {
	profit := <-c
	log.Printf("Profit - %f: %s\n", profit, symbol)
	wg.Done()
}

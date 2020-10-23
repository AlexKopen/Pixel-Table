package main

import (
	"log"
	"sync"
)

// :)
var wg sync.WaitGroup

func main() {
	//	Read stream data for each symbol
	for _, symbol := range Symbols {
		// Process newly acquired stream emissions
		streamGenerationChannel := make(chan []StreamEmission)
		wg.Add(1)
		go generateStreamEmissions(streamGenerationChannel, symbol)
		go receiveStreamGenerationOutput(streamGenerationChannel, symbol)
	}

	wg.Wait()
}

func receiveStreamGenerationOutput(c chan []StreamEmission, symbol string) {
	// After receiving the generated streams, send them off for processing
	streams := <-c
	processStreamChannel := make(chan BotState)
	//log.Printf("Streams fetched: %s\n", symbol)
	go processStreamData(processStreamChannel, streams, symbol)
	go receiveProcessStreamDataOutput(processStreamChannel, symbol)
}

func receiveProcessStreamDataOutput(c chan BotState, symbol string) {
	// After processing the streams, mark the execution as complete
	botState := <-c
	//log.Printf("Bot State - %+v: \n", botState)
	log.Printf("Profit - %f: %s\n", botState.Profit, symbol)
	wg.Done()
}

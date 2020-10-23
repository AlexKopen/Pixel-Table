package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

// :)
var wg sync.WaitGroup
var totalProfit float64
var allBotStates []BotState

func main() {
	// Hide date in logs
	log.SetFlags(0)
	//	Read stream data for each symbol
	for _, symbol := range Symbols {
		// Process newly acquired stream emissions
		streamGenerationChannel := make(chan []StreamEmission)
		wg.Add(1)
		go generateStreamEmissions(streamGenerationChannel, symbol, 1603417344657)
		go receiveStreamGenerationOutput(streamGenerationChannel, symbol)
	}

	// Wait for all symbols to process
	wg.Wait()

	//	Output the final stats
	log.Printf("Total profit: %f\n", totalProfit)
	file, _ := json.Marshal(allBotStates)
	_ = ioutil.WriteFile("output.json", file, 0644)
}

func receiveStreamGenerationOutput(c chan []StreamEmission, symbol string) {
	// After receiving the generated streams, send them off for processing
	streams := <-c
	processStreamChannel := make(chan BotState)
	go processStreamData(processStreamChannel, streams, symbol)
	go receiveProcessStreamDataOutput(processStreamChannel)
}

func receiveProcessStreamDataOutput(c chan BotState) {
	// After processing the streams, mark the execution as complete
	defer wg.Done()
	botState := <-c

	// Stats
	totalProfit += botState.Profit
	log.Printf("Profit - %f: %s\n", botState.Profit, botState.Symbol)
	allBotStates = append(allBotStates, botState)
}

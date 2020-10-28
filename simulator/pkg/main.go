package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pixel-table/simulator/models"
	"sync"
)

// :)
var wg sync.WaitGroup
var allBotStates []models.BotState

func main ()  {
	// Listen for config update requests
	http.HandleFunc("/", configUpdate)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func configUpdate(w http.ResponseWriter, r *http.Request){
	// If a POST request is received, run the simulation with the passed in configuration
	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var configuration models.EngineConfiguration
		configErr := json.Unmarshal(reqBody, &configuration)

		if configErr != nil {
			log.Println("Config parse error")
		} else {
			runSimulation(configuration)
		}
	}
}

func runSimulation(configuration models.EngineConfiguration) {
	// Clear bot states
	allBotStates = nil
	//	Read stream data for each symbol
	for _, symbol := range configuration.Symbols {
		// Process newly acquired stream emissions
		streamGenerationChannel := make(chan []models.StreamEmission)
		wg.Add(1)
		go generateStreamEmissions(streamGenerationChannel, symbol, 1603417344657)
		go receiveStreamGenerationOutput(streamGenerationChannel, symbol)
	}

	// Wait for all symbols to process
	wg.Wait()

	//	Output the final stats
	file, _ := json.Marshal(allBotStates)
	_ = ioutil.WriteFile("output.json", file, 0644)
}

func receiveStreamGenerationOutput(c chan []models.StreamEmission, symbol string) {
	// After receiving the generated streams, send them off for processing
	streams := <-c
	processStreamChannel := make(chan models.BotState)
	go processStreamData(processStreamChannel, streams, symbol)
	go receiveProcessStreamDataOutput(processStreamChannel)
}

func receiveProcessStreamDataOutput(c chan models.BotState) {
	// After processing the streams, mark the execution as complete
	defer wg.Done()
	botState := <-c

	// Stats
	log.Printf("%s complete\n", botState.Symbol)
	allBotStates = append(allBotStates, botState)
}

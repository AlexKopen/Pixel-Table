package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//	Read stream data for each symbol
	for _, symbol := range Symbols {
		// Dynamically create a file path
		filePath := fmt.Sprintf("historical-data/%s.json", strings.ToLower(symbol))
		// Read historical stream data
		historicalData, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		// Covert data read from file to StreamEmission struct
		var streamEmissions []StreamEmission
		unmarshalErr := json.Unmarshal(historicalData, &streamEmissions)
		if unmarshalErr != nil {
			fmt.Println(unmarshalErr)
		}

		// Process newly acquired stream emissions
		processStreamData(streamEmissions, symbol)
	}

}

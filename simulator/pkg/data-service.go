package main

import (
	"Pixel-Table/simulator/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func generateStreamEmissions(c chan []models.StreamEmission, symbol string, endTime int64) {
	// Fetch coin data from the Binance API
	url := fmt.Sprintf("https://api.binance.com/api/v3/klines?interval=5m&symbol=%sUSDT&limit=1000&endTime=%d", symbol, endTime)
	resp, apiErr := http.Get(url)
	if apiErr != nil {
		log.Println("API error: ", apiErr)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Unmarshal the data using a custom interface
	var emissionRef models.StreamEmission
	streamEmissions := []interface{}{emissionRef.OpenTime, emissionRef.Open, emissionRef.High, emissionRef.Low, emissionRef.Close, emissionRef.Volume, emissionRef.CloseTime, emissionRef.QuoteAssetVolume, emissionRef.NumberOfTrades, emissionRef.TakerBuyBaseAssetVolume, emissionRef.TakerBuyQuoteAssetVolume, emissionRef.Ignore}

	if unmarshalErr := json.Unmarshal(body, &streamEmissions); unmarshalErr != nil {
		log.Println("Unmarshal error: ", unmarshalErr)
	}

	// Iterate through the stream emissions with a custom interface and map over to a usable struct
	var streamEmissionsConverted []models.StreamEmission

	for _, record := range streamEmissions {
		if rec, ok := record.([]interface{}); ok {
			tempStreamEmission := models.StreamEmission{}
			for key, val := range rec {
				switch key {
				case 0:
					tempStreamEmission.OpenTime = int64(val.(float64))
				case 1:
					tempStreamEmission.Open = val.(string)
				case 2:
					tempStreamEmission.High = val.(string)
				case 3:
					tempStreamEmission.Low = val.(string)
				case 4:
					tempStreamEmission.Close = val.(string)
				case 5:
					tempStreamEmission.Volume = val.(string)
				case 6:
					tempStreamEmission.CloseTime = int64(val.(float64))
				case 7:
					tempStreamEmission.QuoteAssetVolume = val.(string)
				case 8:
					tempStreamEmission.NumberOfTrades = uint(val.(float64))
				case 9:
					tempStreamEmission.TakerBuyBaseAssetVolume = val.(string)
				case 10:
					tempStreamEmission.TakerBuyQuoteAssetVolume = val.(string)
				case 11:
					tempStreamEmission.Ignore = val.(string)
				}
			}

			streamEmissionsConverted = append(streamEmissionsConverted, tempStreamEmission)

		} else {
			fmt.Printf("StreamEmission struct conversion error")
		}
	}

	// Send the generated stream emissions over the channel
	c <- streamEmissionsConverted
}

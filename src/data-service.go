package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchCoinData(symbol string) string {
	url := fmt.Sprintf("https://api.binance.com/api/v3/klines?startTime=1603042578682&interval=3m&symbol=%sUSDT&limit=1000", symbol)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

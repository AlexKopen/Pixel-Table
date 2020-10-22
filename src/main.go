package main

func main() {
	processStreamData(TestStreams)

	// STEP 1
	//	Read historical stream data for each symbol
	//for _, val := range Symbols {
	//	fmt.Println(val)
	//	historicalData, err := ioutil.ReadFile("historical-data/uni.json")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	jsonparser.ArrayEach(historicalData, func(streamEmission []byte, dataType jsonparser.ValueType, offset int, err error) {
	//		fmt.Println(jsonparser.Get(streamEmission))
	//	})

	//fmt.Printf("File contents: %s", historicalData)
	//}

}

package main

import (
	"./mypkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	jsonFile, err := os.Open("pogoda.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened pogoda.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var weather mypkg.StructWeather

	json.Unmarshal(byteValue, &weather)
	fmt.Printf("Results: %v\n", byteValue)
}

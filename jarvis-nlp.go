package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jarvisNLP/network"
)

type settings struct {
	NewNetwork bool
}

func main() {
	mySettings, err := readConfigFile()
	if err != nil {
		fmt.Printf("Error reading config file %v\n", err)
		return
	}

	fmt.Println("Initializing network")
	net := network.SigmoidNeuronNetwork{}
	net.Initialize(mySettings.NewNetwork)

	fmt.Println("Running network")
	net.Process()

	fmt.Println("Destroy network")
	net.Destroy()
}

func readConfigFile() (result *settings, err error) {
	result = &settings{}

	// Open the configuration file
	fmt.Println("Reading Config File")
	fileData, err := ioutil.ReadFile("config.json")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(fileData, result)
	if err != nil {
		return result, err
	}

	return result, nil
}

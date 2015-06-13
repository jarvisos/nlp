/*
  Copyright 2015 W. Max Lees

  This file is part of jarvisos.

  Jarvisos is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  Jarvisos is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with jarvisos.  If not, see <http://www.gnu.org/licenses/>.

  File: jarvis-nlp.go
  Author: W. Max Lees <max.lees@gmail.com>
  Date: 06.12.2015
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/jarvisos/nlp/network"
	"io/ioutil"
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

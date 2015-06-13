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
	"flag"
	"fmt"
	"github.com/jarvisos/nlp/network"
	"github.com/jarvisos/nlp/settings"
)

func main() {
	// Load the settings
	err := settings.Settings.LoadSettings()
	if err != nil {
		fmt.Printf("Error reading config file %v\n", err)
		return
	}

	// Get the command line flags
	newNet := flag.Bool("n", false, "Forces the program to generate a new neural network")
	flag.Parse()

	// Set any settings based on flags
	settings.Settings.NewNetwork = *newNet

	// Initialize the network
	net := network.SigmoidNeuronNetwork{}
	net.Initialize(settings.Settings.NewNetwork)

	// Process a line
	net.Process()

	// Destroy the network
	net.Destroy()
}

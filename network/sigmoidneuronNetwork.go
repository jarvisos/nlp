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

  File: sigmoidNeuronNetwork.go
  Author: W. Max Lees <max.lees@gmail.com>
  Date: 06.12.2015
*/

package network

import (
	"encoding/gob"
	"fmt"
	"github.com/gonum/matrix/mat64"
	"github.com/jarvisos/nlp/settings"
	"os"
)

type SigmoidNeuronNetwork struct {
	Network []sigmoidNeuronLayer
}

func (n *SigmoidNeuronNetwork) Initialize(newNetwork bool) bool {
	fmt.Println("Initializing Neural Network")

	if newNetwork == true {
		n.createNew()
	} else {
		n.loadFromFile()
	}

	return true
}

func (n *SigmoidNeuronNetwork) createNew() {
	fmt.Println("Generating new Neural Network")

	n.Network = make([]sigmoidNeuronLayer, 3) // 3 Layers

	n.Network[0].initLayer(100, 10)
	n.Network[1].initLayer(10, 10)
	n.Network[2].initLayer(10, 3)

	n.Network[0].randomize()
	n.Network[1].randomize()
	n.Network[2].randomize()
}

func (n *SigmoidNeuronNetwork) Process() *mat64.Vector {
	fmt.Println("Processing input")

	// Generate first level of input
	in := mat64.NewVector(40000, nil)

	return in
	//return n.network[2].process(n.network[1].process(n.network[0].process(in)))
}

func (n *SigmoidNeuronNetwork) SetInput(in string) {

}

func (n *SigmoidNeuronNetwork) Destroy() bool {
	fmt.Println("Cleaning up network")

	err := n.saveToFile()

	if err != nil {
		fmt.Printf("Error saving network to file: %v\n", err)
		return false
	}

	return true
}

func (n *SigmoidNeuronNetwork) loadFromFile() error {
	fmt.Println("Loading network from file")

	file, err := os.OpenFile("network.gob", os.O_RDONLY, os.FileMode(0666))
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(file)

	err = dec.Decode(n)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

func (n *SigmoidNeuronNetwork) saveToFile() error {
	fmt.Println("Saving nework to file")

	// Change settings to load saved network
	settings.Settings.NewNetwork = false
	settings.Settings.SaveSettings()

	file, err := os.OpenFile("network.gob", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(file)

	err = enc.Encode(n)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

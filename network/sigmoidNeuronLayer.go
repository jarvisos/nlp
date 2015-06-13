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

  File: sigmoidNeuronLayer.go
  Author: W. Max Lees <max.lees@gmail.com>
  Date: 06.12.2015
*/

package network

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math/rand"
)

type sigmoidNeuronLayer struct {
	Neurons *mat64.Dense
}

func (layer *sigmoidNeuronLayer) initLayer(in int, size int) {
	fmt.Printf("Initializing single layer of size %v and input size %v\n", size, in)
	layer.Neurons = mat64.NewDense(size, in, nil)
}

func (layer *sigmoidNeuronLayer) process(in *mat64.Vector) (result *mat64.Vector) {
	result.MulVec(layer.Neurons, true, in)

	return result
}

func (layer *sigmoidNeuronLayer) randomize() {
	rand.Seed(756)
	r, c := layer.Neurons.Dims()
	for i := 0; i < r; i++ {
		row := layer.Neurons.RowView(i)
		for j := 0; j < c; j++ {
			row.Set(j, 0, rand.Float64())
		}
	}
}

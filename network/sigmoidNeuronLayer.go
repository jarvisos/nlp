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

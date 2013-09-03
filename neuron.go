package neural

import (
	"log"
	"math/rand"
	"time"
)

type Neuron struct {
	weights []float64
}

func NewNeuron(weights []float64) *Neuron {
	return &Neuron{
		weights,
	}
}

func RandomNeuron(numInputs int) *Neuron {
	weights := make([]float64, numInputs)

	rand.Seed(time.Now().UnixNano())
	rand.Float64()

	for i := range weights {
		weights[i] = rand.Float64()
	}

	return NewNeuron(weights)
}

func (neuron *Neuron) Score(inputs []float64) float64 {
	if len(neuron.weights) != len(inputs) {
		log.Fatalf("wrong number of inputs for neuron: expected %v, got %v\n",
			len(neuron.weights), len(inputs))
		return 0
	}

	return dotProduct(inputs, neuron.weights)
}

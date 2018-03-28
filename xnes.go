package opt

import (
	"sync"

	"golang.org/x/exp/rand"
)

// XNES is the Exponential Natural Evolution Strategies optimiser. It is an NES optimiser
// that uses a multinormal search distribution, taking advantage of a  closed form computation
// of the fisher information matrix.
type XNES struct {
	// Generation data
	size        uint
	searchCount uint
	showCount   uint
	scores      []float64
	seeds       []int64

	// Search distribution parameters
	len   uint
	loc   []float64
	scale float64
	shape []float64

	// Search hyperparameters
	rate     float64
	adaptive bool

	// Noise source
	source *rand.Rand

	// Mutex
	*sync.Mutex
}

// NewXNES creates an XNES optimiser over the d-dimensional real numbers, using the provided
// options for the optimiser.
func NewXNES(d uint, o *Options) (s *SNES) {
	scale := make([]float64, d)
	for i := range scale {
		scale[i] = initScale
	}

	s = &SNES{
		size:        o.GenerationSize,
		showCount:   0,
		searchCount: 0,
		scores:      make([]float64, o.GenerationSize),
		seeds:       make([]int64, o.GenerationSize),

		len:   d,
		loc:   make([]float64, d),
		scale: scale,
		//shape:  IDENTITY MATRIX

		rate:     o.LearningRate,
		adaptive: o.Adaptive,
		source:   rand.New(rand.NewSource(uint64(o.RandomSeed))),

		Mutex: &sync.Mutex{},
	}
	return s
}

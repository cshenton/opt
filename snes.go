package opt

import (
	"math"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// SNES is a Separable Natural Evolution Strategies optimiser. It is a search distribution based
// optimizer that uses a diagonal normal distribution for search.
type SNES struct {
	// Generation data
	size        uint
	searchCount uint
	showCount   uint
	scores      []float64
	seeds       []int64

	// Search distribution parameters
	len   uint
	loc   []float64
	scale []float64

	// Search hyperparameters
	rate     float64
	adaptive bool

	// Noise source
	source *rand.Rand

	// Mutex
	*sync.Mutex
}

const initScale = 1e3

// NewSNES creates a SNES optimiser with the provided parameters and starts its run goroutine.
func NewSNES(len, size uint, seed int64, rate float64, adaptive bool) (s *SNES) {
	scale := make([]float64, len)
	for i := range scale {
		scale[i] = initScale
	}

	s = &SNES{
		size:        size,
		showCount:   0,
		searchCount: 0,
		scores:      make([]float64, size),
		seeds:       make([]int64, size),

		len:   len,
		loc:   make([]float64, len),
		scale: scale,

		rate:     rate,
		adaptive: adaptive,
		source:   rand.New(rand.NewSource(uint64(seed))),

		Mutex: &sync.Mutex{},
	}
	return s
}

// Search returns a point and the seed used to draw it from the search distribution.
func (s *SNES) Search() (point []float64, seed int64) {
	// Awful, but what's the canonical way to block in this setting?
	if s.searchCount >= s.size {
		time.Sleep(10 * time.Nanosecond)
	}

	s.Lock()
	seed = s.source.Int63()
	point = make([]float64, s.len)
	noise := s.makeNoise(seed)
	for i := range point {
		point[i] = s.loc[i] + s.scale[i]*noise[i]
	}
	s.searchCount++
	s.Unlock()
	return point, seed
}

// Show updates the search distribution given a score and the seed that achieved it.
func (s *SNES) Show(score float64, seed int64) {
	s.Lock()
	s.scores[s.showCount] = score
	s.seeds[s.showCount] = seed
	s.showCount++

	if s.showCount >= s.size {
		u := utilities(s.scores)
		gradLoc := make([]float64, s.len)
		gradScale := make([]float64, s.len)
		noise := make([]float64, s.len)
		for i := range u {
			noise = s.makeNoise(s.seeds[i])
			for j := range noise {
				gradLoc[j] += u[i] * noise[j]
				gradScale[j] += u[i] * (math.Pow(noise[j], 2) - 1)
			}
		}

		// if s.adaptive { do thing }

		for i := range s.loc {
			s.loc[i] += s.rate * s.scale[i] * gradLoc[i]
			s.scale[i] *= math.Exp(0.5 * s.rate * gradScale[i])
		}

		s.showCount = 0
		s.searchCount = 0
	}
	s.Unlock()
}

// Precision ... write me tomorrow
func (s *SNES) Precision() (p float64) {
	for i := range s.scale {
		p += s.scale[i]
	}
	return p
}

// makeNoise deterministically makes an vector of standard normal noise from the provided seed.
func (s *SNES) makeNoise(seed int64) (noise []float64) {
	noise = make([]float64, s.len)
	src := rand.New(rand.NewSource(uint64(seed)))
	for i := range noise {
		noise[i] = src.NormFloat64()
	}
	return noise
}

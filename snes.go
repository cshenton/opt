package opt

import "math/rand"

// SNES is a Separable Natural Evolution Strategies optimiser. It is a search
// distribution based optimizer that uses a diagonal normal distribution for search.
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
	// ...
	adaptive bool

	// Noise source
	source *rand.Rand

	// Channels for concurrent access
	searchChan chan searchReq
	showChan   chan showReq
	doneChan   chan bool
}

const initScale = 1e3

// NewSNES creates a SNES optimiser and starts its run goroutine.
func NewSNES(len, size uint, seed int64, adaptive bool) (s *SNES) {
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

		len:    len,
		loc:    make([]float64, len),
		scale:  scale,
		source: rand.New(rand.NewSource(seed)),

		searchChan: make(chan searchReq),
		showChan:   make(chan showReq),
		doneChan:   make(chan bool),
	}
	go s.run()
	return s
}

// Search returns a point and the seed used to draw it from the search distribution.
func (s *SNES) Search() (point []float64, seed int64) {
	rc := make(chan searchResp)
	s.searchChan <- searchReq{
		respChan: rc,
	}
	r := <-rc
	return r.point, r.seed
}

// Show updates the search distribution given a score and the seed that achieved it.
func (s *SNES) Show(score float64, seed int64) {
	s.showChan <- showReq{
		score: score,
		seed:  seed,
	}
}

// doSearch conducts Search assuming exclusive data structure access.
func (s *SNES) doSearch() (point []float64, seed int64) {
	seed = s.source.Int63()
	point = s.makePoint(seed)
	s.searchCount++
	return point, seed
}

// doShow conducts Show assuming exclusive data structure access.
func (s *SNES) doShow(score float64, seed int64) {
	s.scores[s.showCount] = score
	s.seeds[s.showCount] = seed
	s.showCount++

	if s.showCount >= s.size {
		_ = utilities(s.scores)
		// compute grads
		// compute update
		// apply update
		if s.adaptive {
			// compute alternative update and test, update LR is necessary
		}
		s.showCount = 0
		s.searchCount = 0
	}
}

// makePoint generates a draw from the search distribution given a seed.
func (s *SNES) makePoint(seed int64) (point []float64) {
	point = make([]float64, s.len)
	src := rand.New(rand.NewSource(seed))
	for i := range point {
		point[i] = s.loc[i] + s.scale[i]*src.NormFloat64()
	}
	return point
}

// run is the inner loop of the optimiser, and provides safe access to search data.
// If a full generation of searches has been allocated it will stop consuming from
// the search channel until that generation has been processed.
func (s *SNES) run() {
	for {
		if s.searchCount <= s.size {
			// If the generation still needs to be allocated
			select {
			case req := <-s.searchChan:
				point, seed := s.doSearch()
				req.respChan <- searchResp{
					point: point,
					seed:  seed,
				}
			case req := <-s.showChan:
				s.doShow(req.score, req.seed)
			case <-s.doneChan:
				break
			}
		} else {
			// If we're just waiting on results
			select {
			case req := <-s.showChan:
				s.doShow(req.score, req.seed)
			case <-s.doneChan:
				break
			}
		}

	}
}

package opt

// SNES is a Separable Natural Evolution Strategies optimiser. It is a search
// distribution based optimizer that uses a diagonal normal distribution for search.
type SNES struct {
	// Search distribution parameters
	loc   []float64
	scale []float64

	// Channels for concurrent access
	searchChan chan searchReq
	showChan   chan showReq
	doneChan   chan bool
}

// NewSNES creates a SNES optimiser and starts its run goroutine.
func NewSNES() (s *SNES) {
	// Creates search, show, and done channels
	// constructs s
	// fires off go s.run()
	// returns s
	return
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
	return
}

// doSearch conducts Search assuming exclusive data structure access.
func (s *SNES) doSearch() (point []float64, seed int64) {
	// generate seed from rand
	// use seed to draw randNorm
	// scale and return point and seed
	// increment generation counter
	return
}

func (s *SNES) doShow(score float64, seed int64) {
	// Add the results to the gen
	// If the generation is complete
	// compute score
	// compute grads
	// compute update
	// if adaptive compute alternative update and test, update LR is necessary
	// apply update
	// reset generation counter
	return
}

// run is the inner loop of the optimiser, and provides safe access to search data.
// If a full generation of searches has been allocated
func (s *SNES) run() {
	for {
		if true {
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

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
	// Sends a search request to search channel
	// Gets response on response channel
	// Will of course block in optimiser is handling a slow update
	return
}

// Show updates the search distribution given a score and the seed that achieved it.
func (s *SNES) Show(score float64, seed int64) {
	// Sends show request to show channel
	// returns
	// we don't block on show, only on search
	return
}

// doSearch conducts Search assuming exclusive data structure access.
func (s *SNES) doSearch() (point []float64, seed int64) {
	return
}

func (s *SNES) doShow(score float64, seed int64) {
	return
}

// run is the inner loop of the optimiser, and provides safe access to search data.
func (s *SNES) run() {
	for {
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
	}
}

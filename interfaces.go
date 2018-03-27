package opt

// ByteSearcher defines a search algorithm over sequences of bytes.
// Most likely a genetic algorithm.
type ByteSearcher interface {
	// Returns a point to evaluate, and the seed used to generate it
	Search() (point []byte, seed int64)
	// Informs the searcher of the score a particular seeded draw achieved
	Show(score float64, seed int64)
}

// Float64Searcher defines a search algorithm on the double precision n-dimensional real numbers.
// Most likely an evolution strategies algorithm.
type Float64Searcher interface {
	// Returns a point to evaluate, and the seed used to generate it
	Search() (point []float64, seed int64)
	// Informs the searcher of the score a particular seeded draw achieved
	Show(score float64, seed int64)
}

// searchReq is the request sent in a Search call.
type searchReq struct {
	respChan chan<- searchResp
}

// searchReq is the response received in a Search call.
type searchResp struct {
	point []float64
	seed  int64
}

// showReq is the request sent is a Show call.
type showReq struct {
	score float64
	seed  int64
}

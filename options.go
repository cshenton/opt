package opt

// Options is a configuration struct which can holds options commond to several of
// the optimisers in the package.
type Options struct {
	Adaptive       bool    // Whether to use an adaptive learning rate, if available
	GenerationSize uint    // Number of score evaluations per gradient update
	LearningRate   float64 // The learning rate
	RandomSeed     uint64  // Used to seed the source used to generate seeds for searches
}

// DefaultOptions is the recommended set of initial options for most optimisers.
var DefaultOptions = &Options{
	Adaptive:       true,
	GenerationSize: 10,
	LearningRate:   0.01,
	RandomSeed:     24601,
}

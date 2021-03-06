# Opt: Scalable Optimisation for Humans
[![Build Status](https://img.shields.io/travis/cshenton/opt.svg)](https://travis-ci.org/cshenton/opt)
[![Coverage](https://img.shields.io/coveralls/github/cshenton/opt.svg)](https://coveralls.io/github/cshenton/opt)
[![Go Report Card](https://goreportcard.com/badge/github.com/cshenton/opt)](https://goreportcard.com/report/github.com/cshenton/opt)


`opt` is an optimisation library designed to make writing scalable optimisation
routines easy. It provides a unified API for optimising within a single thread,
multiple threads, or across machines, meaning the choice of how to distribute work
is up to you.

Just `go get github.com/cshenton/opt` to install.


## Basic Usage

Opt uses [natural evolution strategies](http://www.jmlr.org/papers/volume15/wierstra14a/wierstra14a.pdf)
optimisers under the hood, which enables a simple API. First, `Search()` against an optimiser
to get a test point, then evaluate its score how you see fit, then `Show()` that score with
the test seed to the optimiser. Keep going until you converge.

```go
package main

import (
	"fmt"

	"github.com/cshenton/opt"
	"github.com/cshenton/opt/bench"
)

func main() {
	n := 10000

	op := opt.DefaultOptions
	op.LearningRate = 0.5
	o := opt.NewSNES(10, op)

	for i := 0; i < n; i++ {
		point, seed := o.Search()
		// Minimise the 10-dimensional sphere function
		score := -bench.Sphere(point)
		o.Show(score, seed)
	}

	final, _ := o.Search()
	fmt.Println(final)
}

```

Want to use more workers on the same machine to speed up evaluations? Just do it,
`opt` optimisers are thread safe, and will simply block calls to `Search` and `Show`
when important computations are happening.

See `examples` for some suggestions on how to use `opt`.


## Choice of Optimizers

Right now, `opt` deals with optimisers that work with real-valued inputs.

#### Available
- SNES (separable natural evolution strategies), great for high (>100) dimensional problems.

#### Coming Soon
- xNES (exponential natural evolution strategies), great for tricky low (<100) dimensional problems.
- Adaptive learning rates for XNES, SNES, great for unfamiliar problems.
- Hyperprojection SNES
- Block-free SNES


## ToDo

- Simpler options handling
- xNES
- More expensive benchmark function to demonstrate multithread perf.
- worker queue example
- rpc example
- benchmark runner.
- some introspection / stopping functions for SNES
- Travis, coveralls

## Dependencies

- `golang.org/x/exp/rand` Rob Pike's experimental rand library.
    - <1ns to create a random source vs. 8.5µs in std lib.
    - This probably matters more for smaller optimisation problems.

# Opt: Scalable Optimisation for Humans

`opt` is an optimisation library designed to make writing scalable optimisation
routines easy. It provides a unified API for optimising within a single thread,
multiple threads, or across machines, meaning the choice of how to distribute work
is up to you.

Just `go get github.com/cshenton/opt` to install.


## Basic Usage

Opt uses evolution strategies optimisers under the hood, which enables a simple API.
First, `Search()` against an optimiser to get a test point, then evaluate its score how
you see fit, then `Show()` that score with the test seed to the optimiser. Keep going
until you converge.

```go
package main

import (
        "github.com/cshenton/opt"
        "github.com/cshenton/opt/bench"
)

func main() {
        n := 1000
        o := opt.NewSNES(...)

        for i := 0; i < n; i++ {
                point, seed := o.Search()
                // Minimise the rastrigin function
                score := -bench.Rastrigin(point)
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

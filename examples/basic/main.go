package main

import (
	"fmt"
	"time"

	"github.com/cshenton/opt"
	"github.com/cshenton/opt/bench"
)

func main() {
	n := 20000
	o := opt.NewSNES(10, 10, 42, 0.1, false)

	t := time.Now()
	for i := 0; i < n; i++ {
		point, seed := o.Search()
		score := -bench.Sphere(point)
		o.Show(score, seed)
	}

	final, _ := o.Search()
	score := -bench.Sphere(final)

	dur := time.Now().Sub(t)
	avg := time.Duration(float64(dur) / float64(n))
	fmt.Printf("\nLoss of %v\n", score)
	fmt.Printf("%v iterations in %v, average %v\n\n", n, dur, avg)
}

// 21.85 micro base
// ~18.5 micro without comms overhead
// 16.4 micro explained by rand.New
// leaves ~1.5 microseconds unexplained
// compared to ~400ns from before
// single makeNoise call takes 8 micro seconds, so those are the source of
// Most of which is the rand.New call...
// Okay we need to fix this...
// So between the two rand.New calls that explains most the difference.
// Even the goroutine overhead hurts :(

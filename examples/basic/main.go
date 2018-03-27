package main

import (
	"fmt"
	"time"

	"github.com/cshenton/opt"
	"github.com/cshenton/opt/bench"
)

func main() {
	o := opt.NewSNES(100, 10, 42, 0.1, false)

	t := time.Now()

	p := 1.0
	n := 0

	for p > 1e-10 {
		point, seed := o.Search()
		score := -bench.Sphere(point)
		o.Show(score, seed)

		p = o.Precision()
		n++
	}

	final, _ := o.Search()
	score := -bench.Sphere(final)

	dur := time.Now().Sub(t)
	avg := time.Duration(float64(dur) / float64(n))
	fmt.Printf("\nLoss of %v\n", score)
	fmt.Printf("%v iterations in %v, average %v\n\n", n, dur, avg)
}

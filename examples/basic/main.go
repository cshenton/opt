package main

import (
	"fmt"
	"time"

	"github.com/cshenton/opt"
	"github.com/cshenton/opt/bench"
)

func main() {
	o := opt.NewSNES(10, 10, 42, 0.01, false)
	t := time.Now()
	p := 1.0
	n := 0

	for p > 1e-8 {
		point, seed := o.Search()
		score := -bench.Rastrigin(point)
		// time.Sleep(10 * time.Microsecond)
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

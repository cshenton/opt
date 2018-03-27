package main

import (
	"fmt"
	"time"

	"github.com/cshenton/opt"
	"github.com/cshenton/opt/bench"
)

func main() {
	o := opt.NewSNES(10, 10, 42, 0.01, false)
	w := 8
	p := 1.0
	done := make(chan int)

	t := time.Now()

	// Just start a bunch of workers that send their iteration counts
	// to a channel on termination.
	for i := 0; i < w; i++ {
		go func() {
			n := 0
			for p > 1e-8 {
				point, seed := o.Search()
				score := -bench.Rastrigin(point)
				// time.Sleep(10 * time.Microsecond)
				o.Show(score, seed)

				p = o.Precision()
				n++
			}
			done <- n
		}()
	}

	n := 0
	for i := 0; i < w; i++ {
		n += <-done
	}

	final, _ := o.Search()
	score := -bench.Sphere(final)

	dur := time.Now().Sub(t)
	avg := time.Duration(float64(dur) / float64(n))
	fmt.Printf("\nLoss of %v\n", score)
	fmt.Printf("%v iterations in %v, average %v\n\n", n, dur, avg)
}

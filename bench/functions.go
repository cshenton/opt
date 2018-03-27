// Package bench contains a variety of benchmark optimization functions. See
// wikipedia.org/wiki/Test_functions_for_optimization for more.
package bench

import "math"

// Rastrigin is the n dimensional rastrigin function on z. The minimum
// is the zero vector.
func Rastrigin(z []float64) (f float64) {
	f = 10 * float64(len(z))
	for i := range z {
		f += math.Pow(z[i], 2) - 10*math.Cos(2*math.Pi*z[i])
	}
	return f
}

// Rosenbrock is the n dimensional rosenbrock function evaluated on z. The
// minimum is the ones vector.
func Rosenbrock(z []float64) (f float64) {
	for i := 0; i < len(z)-1; i++ {
		f += 100*math.Pow(z[i+1]-math.Pow(z[i], 2), 2) + math.Pow(1-z[i], 2)
	}
	return f
}

// Sphere is the n dimensional sphere function evaluated on z. The minimum
// is the zero vector.
func Sphere(z []float64) (f float64) {
	for i := range z {
		f += math.Pow(z[i], 2)
	}
	return f
}

package opt

import (
	"testing"
)

func TestUtilities(t *testing.T) {
	// repeat utils because every rank below halfway gets zeroed
	util := []float64{-0.12161282867924957, -0.2, 0.08457025743803287, -0.2, 0.43704257124121665}
	scores := []float64{28, 13, 80000, 3, 1000000}

	u := utilities(scores)

	if len(u) != len(util) {
		t.Errorf("expected length %v, but got %v", len(util), len(u))
	}

	for i := range u {
		if u[i] != util[i] {
			t.Errorf("expected util %v at position %v but got %v", util[i], i, u[i])
		}
	}
}

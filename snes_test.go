package opt

import "testing"

func TestNewSNES(t *testing.T) {
	tt := []struct {
		name     string
		len      uint
		size     uint
		seed     int64
		rate     float64
		adaptive bool
	}{
		{"small", 10, 10, 42, 0.1, false},
		{"large", 100000, 20, 43, 0.001, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSNES(tc.len, tc.size, tc.seed, tc.rate, tc.adaptive)

			if s.len != tc.len {
				t.Errorf("expected len %v, but got %v", tc.len, s.len)
			}
			if s.size != tc.size {
				t.Errorf("expected size %v, but got %v", tc.size, s.size)
			}
			if s.rate != tc.rate {
				t.Errorf("expected rate %v, but got %v", tc.rate, s.rate)
			}
			if s.adaptive != tc.adaptive {
				t.Errorf("expected adaptive %v, but got %v", tc.adaptive, s.adaptive)
			}
		})
	}
}

func TestSNESSearch(t *testing.T) {}

func TestSNESShow(t *testing.T) {}

func TestSNESdoSearch(t *testing.T) {}

func TestSNESdoShow(t *testing.T) {
	s := NewSNES(3, 10, 42, 0.1, false)

	point, seed := s.doSearch()

	if seed == 0 {
		t.Error("seed not set")
	}
	for i := range point {
		if point[i] == 0 {
			t.Errorf("point not set at position %v", i)
		}
	}
	if s.searchCount != 1 {
		t.Error("search count not incremented")
	}
}

func TestSNESmakeNoise(t *testing.T) {
	var seed int64 = 404
	s := NewSNES(10, 10, 42, 0.1, false)

	n1 := s.makeNoise(seed)
	n2 := s.makeNoise(seed)

	for i := range n1 {
		if n1[i] == 0 {
			t.Errorf("noise not set at position %v", i)
		}
		if n1[i] != n2[i] {
			t.Errorf("noise mismatch at position %v: %v and %v", i, n1[i], n2[i])
		}
	}
}

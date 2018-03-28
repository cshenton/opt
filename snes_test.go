package opt

import (
	"testing"
	"time"
)

func TestNewSNES(t *testing.T) {
	tt := []struct {
		name     string
		dim      uint
		size     uint
		seed     uint64
		rate     float64
		adaptive bool
	}{
		{"small", 10, 10, 42, 0.1, false},
		{"large", 100000, 20, 43, 0.001, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := &Options{
				Adaptive:       tc.adaptive,
				GenerationSize: tc.size,
				LearningRate:   tc.rate,
				RandomSeed:     tc.seed,
			}
			s := NewSNES(tc.dim, o)

			if s.len != tc.dim {
				t.Errorf("expected len %v, but got %v", tc.dim, s.len)
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

func TestSNESSearch(t *testing.T) {
	s := NewSNES(3, DefaultOptions)

	point, seed := s.Search()

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

func TestSNESSearchBlock(t *testing.T) {
	s := NewSNES(3, DefaultOptions)
	for i := 0; i < 10; i++ {
		_, _ = s.Search()
	}

	type searchResp struct {
		point []float64
		seed  int64
	}

	result := make(chan searchResp, 1)
	timeout := make(chan bool, 1)

	go func() {
		point, seed := s.Search()
		result <- searchResp{
			point: point,
			seed:  seed,
		}
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		timeout <- true
	}()

	select {
	case r := <-result:
		t.Errorf("s.Search() should have blocked indefinitely, but got %v, %v", r.point, r.seed)
	case <-timeout:
	}
}

func TestSNESShow(t *testing.T) {
	o := DefaultOptions
	o.GenerationSize = 2
	s := NewSNES(3, o)

	s.Show(20, 42)
	if s.showCount != 1 {
		t.Error("showCount did not increment")
	}
	if s.seeds[0] != 42 {
		t.Errorf("expected seed %v at position %v, but got %v", 42, 0, s.seeds[0])
	}
	if s.scores[0] != 20 {
		t.Errorf("expected score %v at position %v, but got %v", 20, 0, s.scores[0])
	}

	s.Show(13.5, 89)
	if s.searchCount != 0 {
		t.Error("searchCount did not reset")
	}
	if s.showCount != 0 {
		t.Error("showCount did not reset")
	}
	for i := range s.loc {
		if s.loc[i] == 0 {
			t.Errorf("loc did not update at position %v", i)
		}
	}
}

func TestSNESmakeNoise(t *testing.T) {
	var seed int64 = 404
	s := NewSNES(10, DefaultOptions)

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

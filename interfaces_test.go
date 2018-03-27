// These tests simply cause compile errors if an optimizer stops fulfilling its interface.
package opt

import "testing"

func TestFloat64SearcherSNES(t *testing.T) {
	s := NewSNES(10, 10, 3, 0.1, false)
	_ = Float64Searcher(s)
}

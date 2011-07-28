package golomb

import (
	"testing"
)

// Test all values provided by the documentation
func TestEarlyGolombGen(t *testing.T) {
	N := []uint64{1, 2,	3,	4,	5,	6,	7,	8,	9,	10,	11,	12,	13,	14,	15}
	GN := []uint64{1,2,	2,	3,	3,	4,	4,	4,	5,	5,	5,	6,	6,	6,	6}
	gen := StartGenerator()
	for i, n := range N {
		e := <-gen
		if n != e.N || GN[i] != e.GN {
			t.Error("Invalid value generation.  Sequence", n, "was", e.GN, "wanted", GN[i])
		}
	}
}

// Test predefined high values from documentation
func TestHighGValues(t *testing.T) {
	gen := StartGenerator()
	for {
		e := <-gen
		if e.N == 1e3 {
			if e.GN != 86 {
				t.Error("Invalid value generation.  Sequence", e.N, "was", e.GN, "wanted", 86)
			}
		}
		if e.N == 1e6 {
			if e.GN != 6137 {
				t.Error("Invalid value generation.  Sequence", e.N, "was", e.GN, "wanted", 6137)
			}
			break
		}
	}
}

// Test Sum results provided by documentation
func TestExampleSum(t *testing.T) {
	s := Sum(1e3)
	if s != 153506976 {
		t.Error("Invalid sum of 1 - 10^3.  Was ", s, "wanted", 153506976)
	}
}
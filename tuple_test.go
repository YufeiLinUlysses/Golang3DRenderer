package main

import (
	"testing"
	"tuple"
)

func testTuple(t *testing.T) {
	v := Vertex{4.3, -4.2, 3.1, 1.0}
	if !v.isPoint() {
		t.Error("Should be a point")
	}
}

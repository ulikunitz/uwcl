package main

import (
	"math/rand"
	"testing"
)

func TestRandomSelect(t *testing.T) {
	rng := rand.New(rand.NewSource(2))
	for j := 0; j < 100; j++ {
		i := randomSelect(rng, 1<<4|1<<1)
		t.Log(i)
		if i != 1 && i != 4 {
			t.Errorf("expected 1 or 4, got %d", i)
		}
	}
}

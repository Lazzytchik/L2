package dev01

import (
	"math"
	"testing"
	"time"
)

func TestAccuracy(t *testing.T) {
	ti := Time{}

	if math.Abs(float64(ti.Now().Sub(time.Now()).Milliseconds())) > 1000 {
		t.Fatalf("Latency is nore than one second")
	}
}

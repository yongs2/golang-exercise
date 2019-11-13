package clockface

import (
	"testing"
	"time"
	"math"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSecondHandAt30Second(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)
	if got != want {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}
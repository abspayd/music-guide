package app

import (
	"testing"
)

func TestDistance(t *testing.T) {
	distance := IntervalDistance(1, 2)
	_, err := IntervalName(distance)
	if err != nil {
		t.Error(err)
	}

	distance = IntervalDistance(2, 1)
	_, err = IntervalName(distance)
	if err != nil {
		t.Error(err)
	}

	distance = IntervalDistance(0, 5)
	_, err = IntervalName(distance)
	if err != nil {
		t.Error(err)
	}
	distance = IntervalDistance(5, 0)
	_, err = IntervalName(distance)
	if err != nil {
		t.Error(err)
	}
}

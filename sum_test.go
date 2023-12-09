package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(1, 2)
	if result != 3 {
		t.Errorf("Expected sum(1, 2) to be 3, but got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := sub(1, 2)
	if result != -1 {
		t.Errorf("Expected sub(1, 2) to be -1, but got %d", result)
	}
}

func TestTimes(t *testing.T) {
	result := times(2, 2)
	if result != 4 {
		t.Errorf("Expected times(2, 2) to be 4, but got %d", result)
	}
}

func TestDiv(t *testing.T) {
	result := div(4, 2)
	if result != 2 {
		t.Errorf("Expected div(4, 2) to be 2, but got %d", result)
	}
}

func TestMod(t *testing.T) {
	result := mod(5, 2)
	if result != 1 {
		t.Errorf("Expected mod(5, 2) to be 1, but got %d", result)
	}
}

func TestSumX(t *testing.T) {
	result := sumX(1, 2, 3, 4, 5)
	if result != 18 {
		t.Errorf("Expected sumX(1, 2, 3, 4, 5) to be 15, but got %d", result)
	}
}

package math_test

import (
	"testing"
	"with-db/math"
)

func TestAdd(t *testing.T) {
    result, err := math.Add(2, 3)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    expected := 5
    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}

func TestSubtract(t *testing.T) {
    result, err := math.Subtract(5, 3)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    expected := 2
    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}

func TestMultiply(t *testing.T) {
    result, err := math.Multiply(2, 3)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    expected := 6
    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}

func TestDivide(t *testing.T) {
    result, err := math.Divide(6, 3)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    expected := 2
    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}

func TestDivideByZero(t *testing.T) {
	_, err := math.Divide(10, 0)

	if err == nil {
		t.Errorf("expected error but got nil")
	}

	expected := "cannot divide by zero"

	if err.Error() != expected {
		t.Errorf("expected error %s, got %s", expected, err.Error())
	}
}
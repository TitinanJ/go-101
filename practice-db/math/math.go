package math

import "fmt"

func Add(a, b int) (int, error) {
	return a + b, nil
}

func Subtract(a, b int) (int, error) {
	return a - b, nil
}

func Multiply(a, b int) (int, error) {
	return a * b, nil
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
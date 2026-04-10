package main

import (
	"errors"
	"fmt"
)

func check(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negetive")
	}
	return "Positive Number", nil
}

func errorWrapping() error {
	_, err := check(-5)
	if err != nil {
		return fmt.Errorf("failure : %w", err)
	}
	return nil
}

func ErrorHandlingDemo() {
	fmt.Println("basic error handling")
	n, err1 := check(-1)
	p, err2 := check(1)
	fmt.Println(n, err1)
	fmt.Println(p, err2)
	fmt.Println("Wrapping of errors")
	err := errorWrapping()
	if err != nil {
		fmt.Println("Wrapped error:", err)
		// Unwrap to see the original
		underlying := errors.Unwrap(err)
		fmt.Println("Underlying error:", underlying)
		// Or check with errors.Is
		if errors.Is(err, errors.New("negetive")) {
			fmt.Println("Cause is 'negetive'")
		}
	}
}

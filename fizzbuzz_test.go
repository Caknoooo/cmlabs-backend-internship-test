package main

import (
	"bytes"
	"testing"
	"fmt"
)

func Check(res int, out *bytes.Buffer) {
	for i := 1; i <= res; i++ {
		if i%15 == 0 {
			fmt.Fprint(out, "FizzBuzz\n")
		} else if i%3 == 0 {
			fmt.Fprint(out, "Fizz\n")
		} else if i%5 == 0 {
			fmt.Fprint(out, "Buzz\n")
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	var buf bytes.Buffer
	Check(100, &buf)
	output := buf.String()
	expectedOutput := "Fizz\nBuzz\nFizz\nFizz\nBuzz\nFizz\nFizzBuzz\nFizz\nBuzz\nFizz\nFizz\nBuzz\nFizz\nFizzBuzz\nFizz\nBuzz\nFizz\nFizz\nBuzz\nFizz\nFizzBuzz\nFizz\nBuzz\nFizz\nFizz\nBuzz\nFizz\nFizzBuzz\nFizz\nBuzz\nFizz\nFizz\nBuzz\nFizz\nFizzBuzz\nFizz\nBuzz\nFizz\nFizz\nBuzz\nFizz\nFizzBuzz\nFizz\nBuzz\nFizz\nFizz\nBuzz\n"
	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}
}

package main

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rhuan")
	expected := "Hello, Rhuan"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

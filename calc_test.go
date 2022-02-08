package main

import (
	"strings"
	"testing"
)

func Test_Constructor(t *testing.T) {
	calc := newCalculator()
	if calc == nil {
		t.Errorf("Expected a new Calculator, got nil")
	}
}

func TestParse(t *testing.T) {
	expected := 36
	actual, err := newCalculator().parse(strings.Split("6 6 *", " "))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
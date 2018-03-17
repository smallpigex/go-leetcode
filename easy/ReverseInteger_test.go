package main

import (
	"testing"
)

func TestReverse_WhenUnits(t *testing.T) {
	units := 0
	atual := reverse(units)

	expected := 0
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenTensIs10_Return1(t *testing.T) {
	tens := 10
	atual := reverse(tens)
	expected := 1
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenTensIs99_Return99(t *testing.T) {
	tens := 99
	atual := reverse(tens)
	expected := 99
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenTensIs75_Return57(t *testing.T) {
	tens := 75
	atual := reverse(tens)
	expected := 57
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenHundredsIs100_Return1(t *testing.T) {
	hundreds := 100
	atual := reverse(hundreds)
	expected := 1
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenHundredsIs999_Return999(t *testing.T) {
	hundreds := 999
	atual := reverse(hundreds)
	expected := 999
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenHundredsIs123_Return321(t *testing.T) {
	hundreds := 999
	atual := reverse(hundreds)
	expected := 999
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenHundredsIs1534236469_Return0(t *testing.T) {
	hundreds := 1534236469
	atual := reverse(hundreds)
	expected := 0
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenHundredsIsN2147483648_Return0(t *testing.T) {
	hundreds := -2147483648
	atual := reverse(hundreds)
	expected := 0
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

func TestReverse_WhenHundredsIsN12_ReturnN21(t *testing.T) {
	hundreds := -12
	atual := reverse(hundreds)
	expected := -21
	if atual != expected {
		t.Error("Reverse Error")
	} else {
		t.Log("success")
	}
}

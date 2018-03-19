package main

import (
	"testing"
)

func TestNumJewelsInStones_WhenOneJewelInStone_Return1(t *testing.T) {
	jewels := "a"
	stones := "abcdefghi"
	num := numJewelsInStones(jewels, stones)
	if num != 1 {
		t.Errorf("Jewels was incorrect, got: %d, want: %d.", num, 1)
	}
}

func TestNumJewelsInStones_WhenOneJewelInStone_Return9(t *testing.T) {
	jewels := "a"
	stones := "aaaaaaaaa"
	expected := 9
	num := numJewelsInStones(jewels, stones)
	if num != expected {
		t.Errorf("Jewels was incorrect, got: %d, want: %d.", num, expected)
	}
}

func TestNumJewelsInStones_WhenTwoTypeJewelInStone_Return3(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"
	expected := 3
	num := numJewelsInStones(jewels, stones)
	if num != expected {
		t.Errorf("Jewels was incorrect, got: %d, want: %d.", num, expected)
	}
}

func TestNumJewelsInStones_WhenThreeTypeJewelInStone_Return3(t *testing.T) {
	jewels := "aAb"
	stones := "aAAbbbb"
	expected := 7
	num := numJewelsInStones(jewels, stones)
	if num != expected {
		t.Errorf("Jewels was incorrect, got: %d, want: %d.", num, expected)
	}
}

func TestNumJewelsInStones_WhenFiftyTwoTypeJewelInStone_Return104(t *testing.T) {
	jewels := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	stones := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	expected := 104
	num := numJewelsInStones(jewels, stones)
	if num != expected {
		t.Errorf("Jewels was incorrect, got: %d, want: %d.", num, expected)
	}
}

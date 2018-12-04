package main

import "testing"

func TestDay2A(t *testing.T){
	answer, err := twoA()

	if err != nil {
		t.Errorf("2A: died", err)
	}

	if answer != 6642 {
		t.Errorf("2A: expected: 6642, actual: %d", answer)
	}
}

func TestDay2ACalculator(t *testing.T){
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	answer := day2ACalculator(input)

	if answer != 12 {
		t.Errorf("2ACalculator: expected: 12, actual: %d", answer)
	}
}

func TestDay2B(t *testing.T){
	answer, err := twoB()

	if err != nil {
		t.Errorf("2A: died", err)
	}

	if answer != "cvqlbidheyujgtrswxmckqnap" {
		t.Errorf("2A: expected: 6642, actual: %s", answer)
	}
}

func TestDay2BCalculator(t *testing.T){
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	answer := day2BCalculator(input)

	if answer != "fgij" {
		t.Errorf("2BCalculator: expected: 'fgij', actual: '%s'", answer)
	}
}

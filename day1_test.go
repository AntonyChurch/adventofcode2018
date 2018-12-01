package main

import "testing"

func TestOneACalculator(t *testing.T){
	test1Input := []int{1, -2, 3, 1}
	test2Input := []int{1, 1, 1}
	test3Input := []int{1, 1, -2}
	test4Input := []int{-1, -2, -3}

	test1 := OneACalculator(test1Input)
	
	if test1 != 3 {
		t.Errorf("oneACalculator: expected: %d, actual %d", 3, test1)
	}

	test2 := OneACalculator(test2Input)
	
	if test2 != 3 {
		t.Errorf("oneACalculator: expected: %d, actual %d", 3, test2)
	}

	test3 := OneACalculator(test3Input)
	
	if test3 != 0 {
		t.Errorf("oneACalculator: expected: %d, actual %d", 0, test3)
	}

	test4 := OneACalculator(test4Input)
	
	if test4 != -6 {
		t.Errorf("oneACalculator: expected: %d, actual %d", -6, test4)
	}
}

package main

import "testing"

func TestOneA(t *testing.T) {
	answer, err := oneA()

	if err != nil {
		t.Errorf("Error running OneA: %e", err)
	}

	if answer != 580 {
		t.Errorf("OneA: expected: 580, actual %d", answer)
	}
}

func TestOneACalculator(t *testing.T){
	test0Input := []int{1, -2, 3, 1}
	test1Input := []int{1, 1, 1}
	test2Input := []int{1, 1, -2}
	test3Input := []int{-1, -2, -3}

	test0 := OneACalculator(test0Input)
	
	if test0 != 3 {
		t.Errorf("oneACalculator: expected: %d, actual %d", 3, test0)
	}

	test1 := OneACalculator(test1Input)
	
	if test1 != 3 {
		t.Errorf("oneACalculator: expected: %d, actual %d", 3, test1)
	}

	test2 := OneACalculator(test2Input)
	
	if test2 != 0 {
		t.Errorf("oneACalculator: expected: %d, actual %d", 0, test2)
	}

	test3 := OneACalculator(test3Input)
	
	if test3 != -6 {
		t.Errorf("oneACalculator: expected: %d, actual %d", -6, test3)
	}
}

func TestOneBCalculator(t *testing.T){
	input0 := []int{1, -2, 3, 1, 1, -2}
	input1 := []int{1, -1}
	input2 := []int{3, 3, 4, -2, -4}
	input3 := []int{-6, 3, 8, 5, -6}
	input4 := []int{7, 7, -2, -7, -4}

	test0 := OneBCalculator(input0)

	if test0 != 2 {
		t.Errorf("oneBCalculator: expected %d, actual %d", 2, test0)
	}

	test1 := OneBCalculator(input1)

	if test1 != 0 {
		t.Errorf("oneBCalculator: expected %d, actual %d", 0, test1)
	}

	test2 := OneBCalculator(input2)

	if test2 != 10 {
		t.Errorf("oneBCalculator: expected %d, actual %d", 10, test2)
	}

	test3 := OneBCalculator(input3)

	if test3 != 5 {
		t.Errorf("oneBCalculator: expected %d, actual %d", 5, test3)
	}

	test4 := OneBCalculator(input4)

	if test4 != 14 {
		t.Errorf("oneBCalculator: expected %d, actual %d", 14, test4)
	}
}

func TestOneB(t *testing.T){
	answer, err := oneB()

	if err != nil {
		t.Errorf("Error running OneA: %e", err)
	}

	if answer != 81972 {
		t.Errorf("OneB: expected: 81972, actual %d", answer)
	}
}

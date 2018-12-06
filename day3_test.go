package main

import "testing"

func TestThreeA(t *testing.T){
}

func TestThreeACalculator(t *testing.T){
	claims := []claim{
		claim{id: 1, x: 1, y: 3, width: 4, height: 4},
		claim{id: 1, x: 3, y: 1, width: 4, height: 4},
		claim{id: 1, x: 5, y: 5, width: 2, height: 2},
	}

	answer := threeACalculator(8, 8, claims)

	if answer != 4 {
		t.Errorf("ThreeACalculator: expected: 4, actual: %d", answer)
	} 
}

func TestGetClaimFromString(t *testing.T){
	input := "#1 @ 387,801: 11x22"

	cla := getClaimFromString(input)

	if cla.id != 1 {
		t.Errorf("GetClaimFromString claim.id: expected: 1, actual: %d", cla.id)
	}

	if cla.x != 387 {
		t.Errorf("GetClaimFromString claim.x: expected: 387, actual: %d", cla.x)
	}

	if cla.y != 801 {
		t.Errorf("GetClaimFromString claim.y: expected: 801, actual: %d", cla.y)
	}

	if cla.width != 11 {
		t.Errorf("GetClaimFromString claim.width: expected: 11, actual: %d", cla.width)
	}

	if cla.height != 22 {
		t.Errorf("GetClaimFromString claim.height: expected: 22, actual: %d", cla.height)
	}
}

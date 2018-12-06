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
	input1 := "#1 @ 387,801: 11x22"
	input2 := "#10 @ 77,274: 26x21"

	cla := getClaimFromString(input1)

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

	cla2 := getClaimFromString(input2)

	if cla2.id != 10 {
		t.Errorf("GetClaimFromString claim.id: expected: 10, actual: %d", cla2.id)
	}

	if cla2.x != 77 {
		t.Errorf("GetClaimFromString claim.x: expected: 77, actual: %d", cla2.x)
	}

	if cla2.y != 274 {
		t.Errorf("GetClaimFromString claim.y: expected: 274, actual: %d", cla2.y)
	}

	if cla2.width != 26 {
		t.Errorf("GetClaimFromString claim.width: expected: 26, actual: %d", cla2.width)
	}

	if cla2.height != 21 {
		t.Errorf("GetClaimFromString claim.height: expected: 21, actual: %d", cla2.height)
	}
}

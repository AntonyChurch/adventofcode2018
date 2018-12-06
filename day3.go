package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id int
	x int
	y int
	width int
	height int
}

func threeA() (int, error) {
	claims, err := getClaimsFromFile("./day3_input")

	if err != nil {
		return 0, err
	}

	return threeACalculator(1000, 1000, claims), nil
}

func threeACalculator(fabricWidth int, fabricHeight int, claims []claim) int {
	fabric := make([][]int, fabricHeight)
	for x := 0; x < fabricHeight; x++ {
		fabric[x] = make([]int, fabricWidth)
	}
	
	for _, claim := range claims {
		for x := claim.x; x <= claim.x + claim.width - 1; x++ {
			for y := claim.y; y <= claim.y + claim.height - 1; y++ {
				if fabric[x][y] == 0 {
					fabric[x][y] = claim.id
				} else {
					fabric[x][y] = -1
				}
			}
		}
	}

	tally := 0
	for x := 0; x < fabricWidth; x++ {
		for y := 0; y < fabricHeight; y++ {
			if fabric[x][y] == -1 {	
				tally++
			}
		}
	}
	
	return tally
}

func getClaimsFromFile(filePath string) ([]claim, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var claims []claim
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cla := getClaimFromString(scanner.Text())

		claims = append(claims, cla)
	}
	
	return claims, nil
}

func getClaimFromString(str string) claim {
	split := strings.Split(str, " ")

	clm := claim{}

	clm.id, _ = strconv.Atoi(split[0][1:2])
	clm.x, _ = strconv.Atoi(split[2][0:3])
	clm.y, _ = strconv.Atoi(split[2][4:7])
	clm.width, _ = strconv.Atoi(split[3][0:2])
	clm.height, _ = strconv.Atoi(split[3][3:5])

	return clm
}

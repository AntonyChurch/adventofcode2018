package main

import(
	"bufio"
	"os"
	"strconv"
)

func oneA() (int, error) {
	file, err := os.Open("./day1_input")

	if err != nil {
		return 0, err
	}
	defer file.Close()
	
	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		ints = append(ints, i)
	}
	
	return OneACalculator(ints), nil
}

func OneACalculator(ints []int) int {
	total := 0
	for _, i := range ints {
		total += i
	}

	return total
}

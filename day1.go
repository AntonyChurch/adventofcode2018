package main

import(
	"bufio"
	"os"
	"strconv"
)

func oneA() (int, error) {
	ints, err := GetIntsFromFile("./day1_input")

	if err != nil {
		return 0, err
	}
	
	return OneACalculator(ints), nil
}

func oneB() (int, error) {
	ints, err := GetIntsFromFile("./day1_input")

	if err != nil {
		return 0, err
	}
	
	return OneBCalculator(ints), nil
}

func OneACalculator(ints []int) int {
	total := 0
	for _, i := range ints {
		total += i
	}

	return total
}

func OneBCalculator(ints []int) int {
	var frequencies []int
	frequencies = append(frequencies, 0)
	total := 0

	for {
		for _, i := range ints {
			total += i

			if frequencyAlreadyHit(frequencies, total) {
				return total
			}

			frequencies = append(frequencies, total)
		}
	}
	
	return 0
}

func frequencyAlreadyHit(frequencies []int, frequency int) bool {
	for _, i := range frequencies {
		if i == frequency {
			return true
		}
	}

	return false
}

func GetIntsFromFile(filename string) ([]int, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
		
	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}

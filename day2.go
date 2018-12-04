package main

import(
	"bufio"
	"os"
)

func twoA() (int, error){
	strings, err := getStringsFromFile("./day2_input")

	if err != nil {
		return 0, err
	}
	
	return day2ACalculator(strings), nil
}

func day2ACalculator(strings []string) int {
	// iterate string, make map with char and count
	twos := 0
	threes := 0
	
	for _, str := range strings {
		characters := make(map[rune]int)
		
		for _, ch := range str {
			characters[ch]++
		}

		twoDone := false
		threeDone := false
		for key, _ := range characters {
			if characters[key] == 2 {
				if !twoDone {
					twoDone = true
					twos++
				}
			}

			if characters[key] == 3 {
				if !threeDone {
					threeDone = true
					threes++
				}
			}
		}
	}
	
	return twos * threes
}

func twoB() (string, error){
	strings, err := getStringsFromFile("./day2_input")

	if err != nil {
		return "", err
	}
	
	return day2BCalculator(strings), nil
}


func day2BCalculator(strings []string) string {
	for _, str1 := range strings {
		for _, str2 := range strings {
			diffs := 0
			diffPlace := -1
			for i := 0; i < len(str1); i++ {
				if str1[i] != str2[i] {
					diffs++
					diffPlace = i
				}
			}

			if diffs == 1 {
				strReturn := ""
				for i := 0; i < len(str1); i++ {
					if i != diffPlace {
						strReturn += string(str1[i])
					}
				}
				return strReturn
			}
		}
	}
	
	return ""
}

func getStringsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings, nil
}

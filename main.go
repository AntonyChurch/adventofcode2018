package main

import "fmt"

func main(){
	a1, err := oneA()

	if err != nil {
		fmt.Errorf("1A died", err)
	}
	
	fmt.Printf("1A. %d\n", a1)

	b1, err := oneB()

	if err != nil {
		fmt.Errorf("1B died", err)
	}

	fmt.Printf("1B. %d\n", b1)

	a2, err := twoA()

	if err != nil {
		fmt.Errorf("2A died", err)
	}

	fmt.Printf("2A. %d\n", a2)

	b2, err := twoB()

	if err != nil {
		fmt.Errorf("2A died", err)
	}

	fmt.Printf("2B. %s\n", b2)
}

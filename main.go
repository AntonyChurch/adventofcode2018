package main

import "fmt"

func main(){
	a1, err := oneA()

	if err != nil {
		fmt.Errorf("1A died", err)
	}
	
	fmt.Printf("1A. %d\n", a1)
}
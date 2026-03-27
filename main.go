package main

import (
	"fmt"
)

func main() {
	for {
		var thon1 int
		var thon2 int
		fmt.Print("input 1: ")
		fmt.Scanln(&thon1)
		fmt.Print("input 2: ")
		fmt.Scanln(&thon2)
		var OPERATIONS string
		fmt.Println("OPERATION\n", "<1> addition\n", "<2> substraction\n", "<3> multiplication\n", "<4> division")
		fmt.Println("=> selection operation..")
		fmt.Scanln(&OPERATIONS)
		if OPERATIONS == "1" {
			fmt.Println("Result: ", thon1+thon2)
		}
		if OPERATIONS == "2" {
			fmt.Println("Result: ", thon1-thon2)
		}
		if OPERATIONS == "3" {
			fmt.Println("Result: ", thon1*thon2)
		}
		if OPERATIONS == "4" {
			if thon2 == 0 {
				fmt.Println("error..", thon1, "not divisible by 0")
			}
			continue

			if thon1%thon2 != 0 {
				fmt.Println("Result:", "opps..", thon1, "is not divisible by", thon2)
				fmt.Println("try again..")
			} else {
				fmt.Println("Result: ", thon1/thon2)
			}
		}

	}
}

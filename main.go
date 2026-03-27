package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var num1 string
		var num2 string
		fmt.Print("Enter first num: ")
		fmt.Scanln(&num1)
		val1, err := strconv.Atoi(num1)

		if err != nil {
			fmt.Println("Error..need only digit")
			continue
		}

		var operation string
		fmt.Println("> 1. Addition")
		fmt.Println("> 2. Multiplication")
		fmt.Println("> 3. Subtraction")
		fmt.Println("> 4. Division")
		fmt.Print("selection operation : ")
		fmt.Scanln(&operation)

		fmt.Print("Enter second num: ")
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Println("Error .. need  only digit")
			continue
		}
		switch operation {
		case "1":
			fmt.Println("Result:", val1+val2)
		case "2":
			fmt.Println("Result:", val1*val2)
		case "3":
			fmt.Println("Result:", val1-val2)
		case "4":
			if val2 == 0 {
				fmt.Println("error..", num1, "is not divisible by 0")
				continue
			}
			fmt.Println("Result:", val1/val2)
		}
	}
}

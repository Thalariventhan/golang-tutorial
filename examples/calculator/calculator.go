package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("CALCULATOR")
	fmt.Println("type 'exit' to quit")
	fmt.Print("<value> <opt> <value>\n \n")
	for {
		var a_str, b_str, operator string
		fmt.Scanf("%s %s %s", &a_str, &operator, &b_str)
		
		if a_str == "" {
			continue
		}
		
		if a_str == "exit" {
			break
		}

		a, err := strconv.ParseFloat(a_str, 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		
		b, err := strconv.ParseFloat(b_str, 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		var result float64
		switch operator {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				fmt.Println("Error: Division by zero")
				continue
			}
			result = a / b
		case "%":
			result = float64(int64(a) % int64(b))
		default:
			fmt.Println("Error: Unknown operator")
			continue
		}

		fmt.Printf("=> %f\n", result)
	}
}

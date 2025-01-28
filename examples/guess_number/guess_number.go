package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("GUESS NUMBER")
	fmt.Printf("Guess a number between 0 and 99\n")
	fmt.Println("type 'exit' to quit")

	b  := rand.Int63n(99)

	var a_str string
	fmt.Scanf("%s", &a_str)

	if a_str == "exit" {
		return
	}

	a, err := strconv.ParseInt(a_str, 10, 64)
	if err != nil {
		fmt.Println("Invalid input")
	}
	

	for a != b {
		fmt.Println("Wrong 🚫")
		fmt.Scanf("%s", &a_str)
		
		if a_str == "exit" {
			return
		}
		
		a_temp, err := strconv.ParseInt(a_str, 10, 64)
		if err != nil {
			fmt.Println("Invalid input")
		}
		a = a_temp
	}

	fmt.Println("Correct ✅")
}
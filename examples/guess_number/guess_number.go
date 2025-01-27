package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("GUESS NUMBER")
	fmt.Println("type 'exit' to quit")

	b := rand.Intn(99)

	var a int
	fmt.Scanf("%d", &a)

	for a != b {
		fmt.Println("Wrong 🚫")
		fmt.Scanf("%d", &a)
	}

	fmt.Println("Correct 👍🏻")
}
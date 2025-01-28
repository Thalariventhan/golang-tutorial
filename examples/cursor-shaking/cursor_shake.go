package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	// Get the display size
	x, y := robotgo.GetScreenSize()
	fmt.Println("The screen size is", x, y)
	// Moving the mouse in square shape
	for {
		robotgo.Move(200, 200)
		robotgo.Move(200, 300)
		robotgo.Move(300, 300)
		robotgo.Move(300, 200)
		robotgo.Move(200, 200)
	}
}

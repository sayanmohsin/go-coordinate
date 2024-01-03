package main

import (
	"fmt"
	"os"
)

var m = &Movement{
	x: 0,
	y: 0,
	z: 0,
}

func main() {
	var mode int

	fmt.Println("\n--- Select an option by entering the corresponding number ---")
	fmt.Println("1: Move Up")
	fmt.Println("2: Move Right")
	fmt.Println("3: Move Down")
	fmt.Println("4: Move Left")
	fmt.Println("5: Move Forward")
	fmt.Println("6: Move Backward")
	fmt.Println("7: Show Current Coordinates")
	fmt.Println("8: Exit Program")
	fmt.Print("Your selection: ")
	fmt.Scan(&mode)

	switch mode {
	case 1:
		m.moveUp()
		fmt.Println("You moved up.")

	case 2:
		m.moveRight()
		fmt.Println("You moved right.")

	case 3:
		m.moveDown()
		fmt.Println("You moved down.")

	case 4:
		m.moveLeft()
		fmt.Println("You moved left.")

	case 5:
		m.moveForward()
		fmt.Println("You moved forward.")

	case 6:
		m.moveBackward()
		fmt.Println("You moved backward.")

	case 7:
		fmt.Println("Current coordinates are:")
		m.showCoordinates()

	case 8:
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Println("Invalid option. Please enter a number between 1 and 8.")
	}

	main()
}

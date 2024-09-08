package reload

import (
	"fmt"
)

func Start_Menu(Args []string) bool {
	var start string

	system := "linux" // or can be "linux" based on the system

	fmt.Println("***************************************")
	fmt.Println("*             Go Reloaded             *")
	fmt.Println("***************************************")
	fmt.Printf(" [1] Old file: %v\n", Args[1])
	fmt.Printf(" [2] New file: %v\n", Args[2])
	fmt.Println("***************************************")

	for {
		fmt.Println("Do you want to start the process? (Y/N): ")
		fmt.Scan(&start)
		if start == "Y" || start == "y" {
			Clear_Console(system)
			return true
		} else if start == "N" || start == "n" {
			Clear_Console(system)
			return false
		} else {
			fmt.Println("Please enter a valid choice Y or N to continue.")
		}
	}
}

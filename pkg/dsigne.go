package reload

import (
	"fmt"
)

func Start_Menu(Args []string) bool {
	var start string

	system := "windows" // linux or windows

	fmt.Println("***************************************")
	fmt.Println("*             Go Reloaded             *")
	fmt.Println("***************************************")
	fmt.Printf(" [1] old file : %v\n", Args[1])
	fmt.Printf(" [2] new file : %v\n", Args[2])
	fmt.Println("***************************************")

	for {
		fmt.Println("You want to start the process Y/N : ")
		fmt.Scan(&start)
		if (start == "Y") || (start == "y") {
			Clear_Console(system)
			return true
		} else if (start == "N") || (start == "n") {
			Clear_Console(system)
			return false
		} else {
			fmt.Println("Please Enter the correct chois Y or N to continue")
		}
	}
}

func Modifications_Menu() int {
	var choice int
	system := "windows" // can be "linux" or "windows"
	Clear_Console(system)

	fmt.Println("****************************************")
	fmt.Println("*             Text Modifier            *")
	fmt.Println("****************************************")
	fmt.Println(" [1] Replace hexadecimal numbers (hex) with decimal")
	fmt.Println(" [2] Replace binary numbers (bin) with decimal")
	fmt.Println(" [3] Convert words to uppercase (up)")
	fmt.Println(" [4] Convert words to lowercase (low)")
	fmt.Println(" [5] Capitalize the first letter of words (cap)")
	fmt.Println(" [6] Apply (low), (up), (cap) to a specified number of words")
	fmt.Println(" [7] Adjust spaces between punctuation and words")
	fmt.Println(" [8] Adjust quotes (' ') around words")
	fmt.Println(" [9] Replace 'a' with 'an' if the next word starts with a vowel or 'h'")
	fmt.Println(" [10] Exit")
	fmt.Println("****************************************")

	for {
		fmt.Println("Please select the modification you want to apply (1-10): ")
		fmt.Scan(&choice)
		if choice >= 1 && choice <= 10 {
			Clear_Console(system)
			return choice
		} else {
			fmt.Println("Invalid choice. Please enter a number between 1 and 10.")
		}
	}
}

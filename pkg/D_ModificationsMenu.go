package reload

import (
	"fmt"
)

func Modifications_Menu(Line, total int) int {
	var choice int
	system := "linux" // or can be "linux" based on the system
	Clear_Console(system)

	fmt.Println("****************************************")
	fmt.Println("*             Text Modifier            *")
	fmt.Println("****************************************")
	fmt.Printf("Line [%v/%v]\n", Line, total)
	fmt.Println(" [0] Apply all modifications")
	fmt.Println(" [1] Replace hexadecimal numbers (hex) with decimal")
	fmt.Println(" [2] Replace binary numbers (bin) with decimal")
	fmt.Println(" [3] Convert words to uppercase (up)")
	fmt.Println(" [4] Convert words to lowercase (low)")
	fmt.Println(" [5] Capitalize the first letter of words (cap)")
	fmt.Println(" [6] Apply (low), (up), (cap) to a specified number of words")
	fmt.Println(" [7] Adjust spaces between punctuation and words")
	fmt.Println(" [8] Adjust single quotes (' ') around words")
	fmt.Println(" [9] Replace 'a' with 'an' if the next word starts with a vowel or 'h'")
	fmt.Println(" [10] Exit")
	fmt.Println("****************************************")

	for {
		fmt.Println("Please select the modification you want to apply (0-10): ")
		fmt.Scan(&choice)
		if choice >= 0 && choice <= 10 {
			Clear_Console(system)
			return choice
		} else {
			fmt.Println("Invalid choice. Please enter a number between 0 and 10.")
		}
	}
}

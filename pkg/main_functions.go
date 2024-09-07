package reload

import (
	"errors"
	"fmt"
	"os"
)

func Read_File() (string, error) {
	var err error
	Args := os.Args
	if !isValidArg(Args) {
		err = errors.New("invalid Args ")
		return "", err
	}
	file, err := os.ReadFile(Args[1])
	if err != nil {
		err = errors.New("File not Found " + Args[1])
		return "", err
	}

	Start := Start_Menu(Args)

	if Start {
		return string(file), err
	} else {
		err = errors.New("you are cansled the prosce")
		return "", err
	}
}

func Modifications(Data string) string {
	var err error

	for {
		choice := Modifications_Menu()

		switch choice {
		case 1:
			fmt.Println("You selected: Replace hexadecimal numbers (hex) with decimal.")
			Data, err = Convert_To(Data, "(hex)")
			if err != nil {
				return string(err.Error())
			}
		case 2:
			fmt.Println("You selected: Replace binary numbers (bin) with decimal.")
			Data, err = Convert_To(Data, "(bin)")
			if err != nil {
				return string(err.Error())
			}
		case 3:
			fmt.Println("You selected: Convert words to uppercase (up).")
			// Call the function to convert words to uppercase
		case 4:
			fmt.Println("You selected: Convert words to lowercase (low).")
			// Call the function to convert words to lowercase
		case 5:
			fmt.Println("You selected: Capitalize the first letter of words (cap).")
			// Call the function to capitalize words
		case 6:
			fmt.Println("You selected: Apply (low), (up), (cap) to a specified number of words.")
			// Call the function to apply transformations to a specified number of words
		case 7:
			fmt.Println("You selected: Adjust spaces between punctuation and words.")
			Data = Punctuations(SplitPunctuations(Data))
		case 8:
			fmt.Println("You selected: Adjust quotes (' ') around words.")
			// Call the function to adjust quotes
		case 9:
			fmt.Println("You selected: Replace 'a' with 'an' if the next word starts with a vowel or 'h'.")
			// Call the function to replace 'a' with 'an'
		case 10:
			fmt.Println("Exiting program.")
			return Data
		}
	}
}

func Write_File(Data string) {
	os.WriteFile(os.Args[2], []byte(Data), 0o644)
}

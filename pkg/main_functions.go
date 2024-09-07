package reload

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Read_File() (string, error) {
	Args := os.Args
	if !isValidArg(Args) {
		return "", errors.New("invalid Args")
	}

	file, err := os.Open(Args[1])
	if err != nil {
		return "", errors.New("File not Found: " + Args[1])
	}
	defer file.Close()

	// Start the process based on Start_Menu's result
	Start := Start_Menu(Args)
	if !Start {
		return "", errors.New("Exiting program.")
	}

	// Using bufio.Scanner to read file line by line
	var fileContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileContent += line + "\n"
	}
	fileContent = fileContent[:len(fileContent)-1]
	// Checking for scanning errors
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fileContent, nil
}

func Modifications(Data string, Line, total int) (string, error) {
	var err error
	all := 0
	for {
		var choice int

		if all == 0 {
			choice = Modifications_Menu(Line, total)
			if choice == 0 {
				all++
			}
		}
		switch {
		case choice == 1 || all == 1:
			//("You selected: Replace hexadecimal numbers (hex) with decimal.")
			Data, err = Convert_To(Data, "(hex)")
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 2 || all == 2:
			//("You selected: Replace binary numbers (bin) with decimal.")
			Data, err = Convert_To(Data, "(bin)")
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 3 || all == 3:
			//("You selected: Convert words to uppercase (up).")
			Data, err = Low_Up_Cap(Data, "(up)")
			if all != 0 {
				all++
			}
		case choice == 4 || all == 4:
			//("You selected: Convert words to lowercase (low).")
			Data, err = Low_Up_Cap(Data, "(low)")
			if all != 0 {
				all++
			}
		case choice == 5 || all == 5:
			//("You selected: Capitalize the first letter of words (cap).")
			Data, err = Low_Up_Cap(Data, "(cap)")
			if all != 0 {
				all++
			}
		case choice == 6 || all == 6:
			//("You selected: Apply (low), (up), (cap) to a specified number of words.")
			Data, err = Low_Up_Cap_Specified(Data)
			if all != 0 {
				all++
			}
		case choice == 7 || all == 7:
			//("You selected: Adjust spaces between punctuation and words.")
			Data = Punctuations(SplitPunctuations(Data))
			if all != 0 {
				all++
			}
		case choice == 8 || all == 8:
			//("You selected: Adjust quotes (' ') around words.")
			Data = Single_Cote(Data)
			if all != 0 {
				all++
			}
		case choice == 9 || all == 9:
			//("You selected: Replace 'a' with 'an' if the next word starts with a vowel or 'h'.")
			// Call the function to replace 'a' with 'an'

			if all != 0 {
				all++
			}
		case choice == 10 || all == 10:
			fmt.Println("Exiting program.")
			return Data, err
		}
	}
}

func Write_File(Data string) error {
	file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Split the Data into lines and write them one by one
	lines := strings.Split(Data, "\n") // Assume this splits the content by lines (or use strings.Split if needed)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	// Flush any remaining buffered data to the file
	return writer.Flush()
}

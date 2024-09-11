package reload

import (
	"fmt"
	"strings"
)

func Modifications_Control(Data string) (string, error) {
	var err error
	Data_slice := strings.Split(Data, "\n")
	New_Data := ""
	for i := 0; i < len(Data_slice); i++ {
		Data, err = Modifications(Data_slice[i], i+1, len(Data_slice))
		if err != nil {
			return "", fmt.Errorf("error Line %d : %s", i+1, err.Error())
		}
		New_Data += Data + "\n"
	}
	return New_Data, err
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
		case choice == 1 || all == 1 || choice == 2 || all == 2:
			//("You selected: Replace hexadecimal numbers (hex) with decimal.") // 1
			//("You selected: Replace binary numbers (bin) with decimal.") // 2
			Data, err = Convert_To(Data)
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 3 || all == 3:
			//("You selected: Convert words to uppercase (up).")
			Data, err = Low_Up_Cap_Specified(Data)
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 4 || all == 4:
			//("You selected: Convert words to lowercase (low).")
			Data, err = Low_Up_Cap_Specified(Data)
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 5 || all == 5:
			//("You selected: Capitalize the first letter of words (cap).")
			Data, err = Low_Up_Cap_Specified(Data)
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 6 || all == 6:
			//("You selected: Apply (low), (up), (cap) to a specified number of words.")
			Data, err = Low_Up_Cap_Specified(Data)
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 7 || all == 7:
			//("You selected: Adjust spaces between punctuation and words.")
			Data = Punctuations(SplitPunctuations(Data))
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 8 || all == 8:
			//("You selected: Adjust quotes (' ') around words.")
			Data = Single_Cote(Data)
			if err != nil {
				return "", err
			}
			if all != 0 {
				all++
			}
		case choice == 9 || all == 9:
			//("You selected: Replace 'a' with 'an' if the next word starts with a vowel or 'h'.")
			Data = AtoAn(Data)
			if all != 0 {
				all++
			}
		case choice == 10 || all == 10:
			fmt.Println("Exiting program.")
			return Data, err
		}
	}
}

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
	for {

		choice := Modifications_Menu(Line, total)

		switch {
		case choice == 0:
			//("You selected: Apply All modifications.")
			Data, err = ApplayAllModif(Data)
			if err != nil {
				return "", err
			}
			fmt.Println("Exiting program.")
			return Data, err
		case choice == 1:
			//("You selected: Replace hexadecimal numbers (hex) with decimal.")
			Data, err = Convert_To(Data, flags.hex)
			if err != nil {
				return "", err
			}
		case choice == 2:
			//("You selected: Replace binary numbers (bin) with decimal.")

			Data, err = Convert_To(Data, flags.bin)
			if err != nil {
				return "", err
			}

		case choice == 3:
			//("You selected: Convert words to uppercase (up).")
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(Data, up, !low, !cap)
			if err != nil {
				return "", err
			}
		case choice == 4:
			//("You selected: Convert words to lowercase (low).")
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(Data, !up, low, !cap)
			if err != nil {
				return "", err
			}
		case choice == 5:
			//("You selected: Capitalize the first letter of words (cap).")
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(Data, !up, !low, cap)
			if err != nil {
				return "", err
			}
		case choice == 6:
			//("You selected: Apply (low), (up), (cap) to a specified number of words.")
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(Data, up, low, cap)
			if err != nil {
				return "", err
			}
		case choice == 7:
			//("You selected: Adjust spaces between punctuation and words.")
			Data = Punctuations(SplitPunctuations(Data))
			if err != nil {
				return "", err
			}
		case choice == 8:
			//("You selected: Adjust quotes (' ') around words.")
			Data = Single_Cote(Data)
			if err != nil {
				return "", err
			}
		case choice == 9:
			//("You selected: Replace 'a' with 'an' if the next word starts with a vowel or 'h'.")
			Data = AtoAn(Data)
		case choice == 10:
			fmt.Println("Exiting program.")
			return Data, err
		case choice == 11:
			return Data, err
		}
	}
}

package reload

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func isvalidExt(file_name string) bool {
	for i, char := range file_name {
		if char == '.' {
			if file_name[i:] == ".txt" {
				return true
			}
			break
		}
	}
	return false
}

func isValidArg(args []string) bool {
	if len(args) != 3 {
		return false
	}
	if isvalidExt(args[1]) && isvalidExt(args[2]) {
		return true
	}
	return false
}

func Convert_By_Bas(Number, Bas string) (string, error) {
	var base int
	if Bas == "hex" {
		base = 16
	} else {
		base = 2
	}
	conver, err := strconv.ParseInt(Number, base, 64)

	return strconv.Itoa(int(conver)), err
}

func Convert_To(Data string) (string, error) {
	Data_Slices := strings.Split(Data, " ")
	New_Data_Slices := []string{}
	var err error

	for i, word := range Data_Slices {
		switch {

		case strings.Contains(word, "(hex)") && i > 0:

			if word != "(hex)" {
				err = errors.New("Syntax Error " + word)
				return "", err
			}

			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], "hex")
			if err != nil {
				fmt.Println(err)
				return "", err
			}

		case strings.Contains(word, "(bin)") && i > 1:

			if word != "(bin)" {
				err = errors.New("Syntax Error " + word)
				return "", err
			}

			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], "bin")
			if err != nil {
				fmt.Println(err)
				return "", err
			}

		default:
			New_Data_Slices = append(New_Data_Slices, word)
		}
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}

func Capit(word string) string {
	cap_word := ""
	for i, char := range word {
		if (char >= 'a' && char <= 'z') && i == 0 {
			cap_word += string(char - 32)
		} else {
			cap_word += string(char)
		}
	}
	return cap_word
}

func Split_Low_Up_Cap(str string) []string {
	NewStr := ""
	str += " "
	slisce := []string{}
	bracket := false

	for _, char := range str {
		if char == '(' {
			bracket = true
		} else if char == ')' {
			bracket = false
		}
		if char == ' ' && NewStr != "" && !bracket {
			slisce = append(slisce, NewStr)
			NewStr = ""
		} else if char != ' ' {
			NewStr += string(char)
		}
	}

	return slisce
}

func Low_Up_Cap(Data string) (string, error) {
	Data_Slices := Split_Low_Up_Cap(Data)
	New_Data_Slices := []string{}
	var err error

	for i, word := range Data_Slices {
		switch {

		case strings.Contains(word, "(hex)") && i > 0:

			if word != "(hex)" {
				err = errors.New("Syntax Error " + word)
				return "", err
			}

			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], "hex")
			if err != nil {
				fmt.Println(err)
				return "", err
			}

		case strings.Contains(word, "(bin)") && i > 1:

			if word != "(bin)" {
				err = errors.New("Syntax Error " + word)
				return "", err
			}

			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], "bin")
			if err != nil {
				fmt.Println(err)
				return "", err
			}

		default:
			New_Data_Slices = append(New_Data_Slices, word)
		}
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}

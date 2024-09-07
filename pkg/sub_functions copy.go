package reload

import (
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
		word = strings.ReplaceAll(word, "(hex)", " (hex) ")
		word = strings.ReplaceAll(word, "(bin)", " (bin) ")
		switch {
		case strings.Contains(word, "(hex)") && i > 1:
			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], "hex")
			if err != nil {
				fmt.Println(err)
				return "", err
			}
		case strings.Contains(word, "(bin)") && i > 1:
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

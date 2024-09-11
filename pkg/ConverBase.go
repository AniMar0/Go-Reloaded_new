package reload

import (
	"strconv"
	"strings"
)

func Convert_By_Bas(Number, Bas string) (string, error) {
	var base int
	if Bas == "(hex)" {
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
		if i == 0 && (word == "(hex)" || word == "(bin)") {
			continue
		}
		switch {
		case word == "(hex)":
			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(New_Data_Slices[len(New_Data_Slices)-1], "(hex)")
			if err != nil {
				return "", err
			}
		case word == "(bin)":
			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(New_Data_Slices[len(New_Data_Slices)-1], "(bin)")
			if err != nil {
				return "", err
			}
		default:
			New_Data_Slices = append(New_Data_Slices, word)
		}
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}

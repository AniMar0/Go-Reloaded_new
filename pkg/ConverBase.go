package reload

import (
	"errors"
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

func Convert_To(Data, Bas string) (string, error) {
	Data_Slices := strings.Split(Data, " ")
	New_Data_Slices := []string{}
	var err error

	for i, word := range Data_Slices {

		switch {
		case i == 0 && word == Bas:
			err = errors.New("Syntax Error " + word)
			return "", err
		case word == Bas:
			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], Bas)
			if err != nil {
				return "", err
			}
		}
		word = strings.ReplaceAll(word, Bas, "")
		New_Data_Slices = append(New_Data_Slices, word)
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}

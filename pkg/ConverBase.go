package reload

import (
	"strconv"
	"strings"
)

func Convert_By_Bas(Data []string, Bas string) ([]string, error) {
	var base int

	if flags.hexFlag(Bas) {
		base = 16
	} else if flags.binFlag(Bas) {
		base = 2
	}
	for i := len(Data) - 1; i >= 0; i-- {
		if !flags.checkFlag(Data[i]) {
			conver, err := strconv.ParseInt(Data[i], base, 0)
			Data[i] = strconv.Itoa(int(conver))
			return Data, err
		}
	}
	return Data, nil
}

func Convert_To(Data, Type string) (string, error) {
	Data_Slices := strings.Split(Data, " ")
	Data_Slices = Clean(Data_Slices)
	New_Data_Slices := []string{}
	var err error

	for i, word := range Data_Slices {
		switch {
		case flags.hexFlag(word) && word == Type:
			if i == 0 {
				continue
			}
			New_Data_Slices, err = Convert_By_Bas(New_Data_Slices, flags.hex)
			if err != nil {
				return "", err
			}
		case flags.binFlag(word) && word == Type:
			if i == 0 {
				continue
			}
			New_Data_Slices, err = Convert_By_Bas(New_Data_Slices, flags.bin)
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


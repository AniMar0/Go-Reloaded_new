package reload

import (
	"errors"
	"strconv"
	"strings"
)

func LUC(Data []string, Number int, luc func(string) string) []string {
	for i := len(Data) - 1; i >= len(Data)-Number; i-- {
		if !flags.checkFlag(Data[i]) {
			Data[i] = luc(Data[i])
		} else {
			Number++
		}
	}
	return Data
}

func Low_Up_Cap_Specified(Data string, up, low, cap bool) (string, error) {
	SliceOfData := SplitUpCapLow(Data)
	SliceOfData = Clean(SliceOfData)
	SliceOfData_New := []string{}
	for i, word := range SliceOfData {
		switch {
		case flags.lowFlag(word) && (!up && low && !cap):
			if i == 0 {
				continue
			}
			SliceOfData_New = LUC(SliceOfData_New, 1, Low)
		case flags.upFlag(word) && (up && !low && !cap):
			if i == 0 {
				continue
			}
			SliceOfData_New = LUC(SliceOfData_New, 1, Up)
		case flags.capFlag(word) && (!up && !low && cap):
			if i == 0 {
				continue
			}
			SliceOfData_New = LUC(SliceOfData_New, 1, Capit)

		case flags.lowSpecified(word) && (up && cap && low):
			if i == 0 {
				continue
			}
			isNumber, Number := IsDigit(word)
			if !isNumber {
				SliceOfData_New = append(SliceOfData_New, word)
				continue
			} else if Number > len(SliceOfData_New) {
				return "", errors.New("error : the number of words you want to lower is more than Number of the words before the flag")
			}

			SliceOfData_New = LUC(SliceOfData_New, Number, Low)
		case flags.upSpecified(word) && (up && cap && low):
			if i == 0 {
				continue
			}
			isNumber, Number := IsDigit(word)
			if !isNumber {
				SliceOfData_New = append(SliceOfData_New, word)
				continue
			} else if Number > len(SliceOfData_New) {
				return "", errors.New("error : the number of words you want to lower is more than Number of the words before the flag")
			}

			SliceOfData_New = LUC(SliceOfData_New, Number, Up)
		case flags.capSpecified(word) && (up && cap && low):
			if i == 0 {
				continue
			}
			isNumber, Number := IsDigit(word)
			if !isNumber {
				SliceOfData_New = append(SliceOfData_New, word)
				continue
			} else if Number > len(SliceOfData_New) {
				return "", errors.New("error : the number of words you want to lower is more than Number of the words before the flag")
			}

			SliceOfData_New = LUC(SliceOfData_New, Number, Capit)

		default:
			SliceOfData_New = append(SliceOfData_New, word)

		}
	}

	Data = strings.Join(SliceOfData_New, " ")

	return Data, nil
}

func IsDigit(word string) (bool, int) {
	if flags.lowSpecified(word) || flags.capSpecified(word) {
		word = word[6 : len(word)-1]
	} else {
		word = word[5 : len(word)-1]
	}
	number, err := strconv.Atoi(word)
	if err != nil {
		return false, number
	}
	return true, number
}

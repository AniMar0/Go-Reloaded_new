package reload

import (
	"errors"
	"strings"
)

func Low_Up_Cap_Specified(Data string) (string, error) {
	SliceOfData := SplitUpCapLow(Data)
	SliceOfData_New := []string{}
	for i, word := range SliceOfData {
		if i == 0 && (word == "(low)" || word == "(up)" || word == "(cap)" || strings.HasPrefix(word, "(low, ") || strings.HasPrefix(word, "(up, ") || strings.HasPrefix(word, "(cap, ")) {
			continue
		}
		// println(word+"***")
		switch {
		case word == "(low)":
			SliceOfData_New[len(SliceOfData_New)-1] = Low(SliceOfData_New[len(SliceOfData_New)-1])

		case word == "(up)":
			SliceOfData_New[len(SliceOfData_New)-1] = Up(SliceOfData_New[len(SliceOfData_New)-1])

		case word == "(cap)":
			SliceOfData_New[len(SliceOfData_New)-1] = Capit(SliceOfData_New[len(SliceOfData_New)-1])

		case strings.HasPrefix(word, "(low, ") && strings.HasSuffix(word, ")"):
			SliceOfData_New = append(SliceOfData_New, "zabii")
			isNumber, Number := IsDigit(word)
			if !isNumber {
				SliceOfData_New = append(SliceOfData_New, word)
				continue
			} else if Number > len(SliceOfData_New) {
				return "", errors.New("error : the number of words you want to lower is more than Number of the words before the flag")
			}

			SliceOfData_New = ChangeWord(SliceOfData_New, Low, Number)

		case strings.HasPrefix(word, "(up, ") && strings.HasSuffix(word, ")"):
			isNumber, Number := IsDigit(word)
			if !isNumber {
				SliceOfData_New = append(SliceOfData_New, word)
				continue
			} else if Number > len(SliceOfData_New) {
				return "", errors.New("error : the number of words you want to lower is more than Number of the words before the flag")
			}

			SliceOfData_New = ChangeWord(SliceOfData_New, Up, Number)
		case strings.HasPrefix(word, "(cap, ") && strings.HasSuffix(word, ")"):
			isNumber, Number := IsDigit(word)
			if !isNumber {
				SliceOfData_New = append(SliceOfData_New, word)
				continue
			} else if Number > len(SliceOfData_New) {
				return "", errors.New("error : the number of words you want to lower is more than Number of the words before the flag")
			}

			SliceOfData_New = ChangeWord(SliceOfData_New, Capit, Number)
		default:
			SliceOfData_New = append(SliceOfData_New, word)

		}
	}

	Data = strings.Join(SliceOfData_New, " ")

	return Data, nil
}

func IsDigit(word string) (bool, int) {
	if strings.HasPrefix(word, "(low, ") || strings.HasPrefix(word, "(cap, ") {
		word = word[5 : len(word)-2]
	} else {
		word = word[4 : len(word)-2]
	}
	number := 0
	isNumber := false
	for _, digit := range word {
		if digit >= 0 && digit <= 9 {
			number = number*10 + int(digit-48)
			isNumber = true
		} else {
			return false, 0
		}
	}
	return isNumber, number
}

func ChangeWord(words []string, Modife func(string) string, Number int) []string {
	for i := len(words) - Number; i < len(words); i++ {
		words[i] = Modife(words[i])
	}
	return words
}

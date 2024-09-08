package reload

import (
	"strconv"
	"strings"
)

func Low_Up_Cap_Specified(Data string) (string, error) {
	Data_Slices := Split_Low_Up_Cap(Data)
	New_Data_Slices := []string{}
	var err error
	for i, word := range Data_Slices {
		if i > 0 {
			_, Type, _ := SplitByPrefixSuffix(word, "(", ")")
			Type = "(" + Type + ")"
			switch {
			case strings.HasPrefix(word, "(low,"):

				ModifyWords(New_Data_Slices, Read_number(word), Low)
				word = strings.ReplaceAll(word, Type, "")
				New_Data_Slices = append(New_Data_Slices, word)

			case strings.HasPrefix(word, "(up,"):

				ModifyWords(New_Data_Slices, Read_number(word), Up)
				word = strings.ReplaceAll(word, Type, "")
				New_Data_Slices = append(New_Data_Slices, word)

			case strings.HasPrefix(word, "(cap,"):

				ModifyWords(New_Data_Slices, Read_number(word), Capit)
				word = strings.ReplaceAll(word, Type, "")
				New_Data_Slices = append(New_Data_Slices, word)
			default:
				New_Data_Slices = append(New_Data_Slices, word)
			}
		} else {
			New_Data_Slices = append(New_Data_Slices, word)
		}
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}

func Read_number(content string) int {
	number := 0
	for _, char := range content {
		if char >= '0' && char <= '9' {
			new_number, _ := strconv.Atoi(string(rune(char)))
			number = number*10 + new_number
		}
	}
	if number == 0 {
		return 1
	}
	return number
}

func ModifyWords(content []string, number int, transformFunc func(string) string) {
	for i := len(content) - number; i < len(content); i++ {
		content[i] = transformFunc(content[i])
	}
}

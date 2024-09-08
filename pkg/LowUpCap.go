package reload

import (
	"fmt"
	"strings"
)

func Low_Up_Cap(Data, Type string) (string, error) {
	Data_Slices := Split_Low_Up_Cap(Data)
	New_Data_Slices := []string{}
	var err error

	for i, word := range Data_Slices {
		if strings.Contains(word, Type) && i > 0 {
			switch Type {
			case "(up)":

				New_Data_Slices[len(New_Data_Slices)-1] = Up(Data_Slices[i-1])
				if err != nil {
					fmt.Println(err)
					return "", err
				}
				word = strings.ReplaceAll(word, Type, "")
				New_Data_Slices = append(New_Data_Slices, word)

			case "(low)":

				New_Data_Slices[len(New_Data_Slices)-1] = Low(Data_Slices[i-1])
				if err != nil {
					fmt.Println(err)
					return "", err
				}
				word = strings.ReplaceAll(word, Type, "")
				New_Data_Slices = append(New_Data_Slices, word)

			case "(cap)":

				New_Data_Slices[len(New_Data_Slices)-1] = Capit(Data_Slices[i-1])
				if err != nil {
					fmt.Println(err)
					return "", err
				}
				word = strings.ReplaceAll(word, Type, "")
				New_Data_Slices = append(New_Data_Slices, word)

			}
		} else {
			New_Data_Slices = append(New_Data_Slices, word)
		}
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}



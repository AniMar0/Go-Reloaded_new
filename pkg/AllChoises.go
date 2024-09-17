package reload

import (
	"fmt"
	"strings"
)

func All_Control(Data string) (string, error) {
	var err error
	Data_slice := strings.Split(Data, "\n")
	New_Data := ""
	for i := 0; i < len(Data_slice); i++ {
		Data, err = ApplayAllModif(Data_slice[i])
		if err != nil {
			return "", fmt.Errorf("error Line %d : %s", i+1, err.Error())
		}
		if i < len(Data_slice)-1 {
			New_Data += Data + "\n"
		} else {
			New_Data += Data
		}
	}
	return New_Data, err
}

func All(Data []string) ([]string, []string, string) {
	for i, Word := range Data {
		if flags.isFlag(Word) {
			if i < len(Data)-1 {
				return Data[:i+1], Data[i+1:], Word
			} else {
				return Data, nil, Word
			}
		}
	}
	return Data, nil, "EptyZaki001+"
}

func joinToParts(Part1, Part2 []string) (string, string) {
	part1 := strings.Join(Part1, " ")
	if Part2 == nil {
		return part1, ""
	}
	part2 := strings.Join(Part2, " ")
	return part1, part2
}

func ApplayAllModif(Data string) (string, error) {
	var err error

	for {
		part1, part2 := "", ""
		cleanData := SplitLowUpCap(Data)
		cleanData = Clean(cleanData)

		Part1, Part2, flag := All(cleanData)
		part1, part2 = joinToParts(Part1, Part2)

		switch {
		case flags.hexFlag(flag):
			Data, err = Convert_To(part1, flags.hex)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.binFlag(flag):
			Data, err = Convert_To(part1, flags.bin)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.upFlag(flag):
			Data, err = Low_Up_Cap_Specified(part1)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.lowFlag(flag):
			Data, err = Low_Up_Cap_Specified(part1)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.capFlag(flag):
			Data, err = Low_Up_Cap_Specified(part1)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.capSpecified(flag) || flags.upSpecified(flag) || flags.lowSpecified(flag):
			Data, err = Low_Up_Cap_Specified(part1)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case !flags.isFlag(flag):
			Data = part1
			Data = Punctuations(SplitPunctuations(Data))
			Data = Single_Cote(Data)
			Data = AtoAn(Data)
			return Data, err
		}
	}
}

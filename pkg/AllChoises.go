package reload

import (
	"strings"
)

func All(Data []string) (int, string) {
	for i, Word := range Data {
		if flags.isFlag(Word) {
			if i < len(Data)-1 {
				return i + 1, Word
			} else {
				return i, Word
			}
		}
	}
	return 0, ""
}

func joinToParts(Data []string, index int) (string, string) {
	parte1 := strings.Join(Data[:index], " ")
	parte2 := strings.Join(Data[index:], " ")
	return parte1, parte2
}

func ApplayAllModif(Data string) (string, error) {
	var err error
	part1, part2 := "", ""
	for {
		cleanData := SplitUpCapLow(Data)
		cleanData = Clean(cleanData)
		index, flag := All(cleanData)
		if index != 0 && flag != "" {
			part1, part2 = joinToParts(cleanData, index)
		}

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
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(part1, up, !low, !cap)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.lowFlag(flag):
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(part1, !up, low, !cap)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.capFlag(flag):
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(part1, !up, !low, cap)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case flags.capSpecified(flag) || flags.upSpecified(flag) || flags.lowSpecified(flag):
			up, low, cap := true, true, true
			Data, err = Low_Up_Cap_Specified(Data, up, low, cap)
			Data += " " + part2
			if err != nil {
				return "", err
			}
		case !flags.isFlag(flag):
			Data = Data[:len(Data)-len(part2)]
			Data = Punctuations(SplitPunctuations(Data))
			if err != nil {
				return "", err
			}
			Data = Single_Cote(Data)
			Data = AtoAn(Data)

			return Data, err
		}
	}
}

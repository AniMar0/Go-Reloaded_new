package reload

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func isvalidExt(file_name string) bool {
	for i, char := range file_name {
		if char == '.' && i != 0 {
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
		if strings.Contains(word, Bas) && i > 0 {

			/*
				if word != Bas {
					err = errors.New("Syntax Error " + word)
					return "", err
				}
			*/

			New_Data_Slices[len(New_Data_Slices)-1], err = Convert_By_Bas(Data_Slices[i-1], Bas)
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			word = strings.ReplaceAll(word, Bas, "")
			New_Data_Slices = append(New_Data_Slices, word)
		} else {
			New_Data_Slices = append(New_Data_Slices, word)
		}
	}

	Data = strings.Join(New_Data_Slices, " ")
	return Data, err
}

func Capit(word string) string {
	cap_word := ""
	if len(word) == 0 || isAlpha(rune(word[0])) {
		// return "", errors.New("You can't capitalize the word: " + word)
		cap_word = strings.ToUpper(string(word[0])) + word[1:]
	}

	return cap_word
}

func Split_Low_Up_Cap(str string) []string {
	NewStr := ""
	str += " "
	slisce := []string{}
	bracket := false

	for _, char := range str {
		if char == '(' {
			bracket = true
		} else if char == ')' {
			bracket = false
		}
		if char == ' ' && NewStr != "" && !bracket {
			slisce = append(slisce, NewStr)
			NewStr = ""
		} else if char != ' ' {
			NewStr += string(char)
		}
	}

	return slisce
}

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

// 6
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

func ModifyWords(content []string, number int, transformFunc func(string) string) {
	for i := len(content) - number; i < len(content); i++ {
		content[i] = transformFunc(content[i])
	}
}

func Low(word string) string {
	return strings.ToLower(word)
}

func Up(word string) string {
	return strings.ToUpper(word)
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

func SplitByPrefixSuffix(s, prefix, suffix string) (string, string, string) {
	startIndex := strings.Index(s, prefix)
	if startIndex == -1 {
		return s, "", ""
	}
	startIndex += len(prefix)

	endIndex := strings.LastIndex(s, suffix)
	if endIndex == -1 || endIndex < startIndex {
		return s, "", ""
	}

	before := s[:startIndex-len(prefix)]
	middle := s[startIndex:endIndex]
	after := s[endIndex+len(suffix):]

	return before, middle, after
}

// ----

func Clear_Console(system string) {
	// For Unix/Linux
	if system == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	// For Windows
	if system == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Punctuations(content []string) string {
	New_content := ""
	for i, w := range content {
		isAllPunc := true // Reset this flag for each word
		for _, c := range w {
			if c != '.' && c != ',' && c != '!' && c != '?' && c != ':' && c != ';' {
				isAllPunc = false
				break
			}
		}
		if isAllPunc {
			New_content += w // No space before punctuation marks
		} else {
			if i > 0 { // Add a space before non-punctuation words only if it's not the first word
				New_content += " "
			}
			New_content += w // Add non-punctuation words
		}
	}
	return New_content
}

func SplitPunctuations(s string) []string {
	var slice []string
	var torf []bool
	var result []string
	wordStart := -1
	inWord := false

	for i, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			if inWord {
				slice = append(slice, s[wordStart:i])
				inWord = false
			}
		} else {
			if !inWord {
				wordStart = i
				inWord = true
			}
		}
	}
	if inWord {
		slice = append(slice, s[wordStart:])
	}

	for _, w := range slice {
		if len(w) > 0 && strings.ContainsAny(string(w[0]), ".,?!:;") {
			torf = append(torf, true)
		} else {
			torf = append(torf, false)
		}
	}

	for i, w := range slice {
		if !torf[i] {
			result = append(result, w)
		} else {
			j := 0
			for j < len(w) && strings.ContainsAny(string(w[j]), ".,?!:;") {
				j++
			}
			if j > 0 {
				firsthalf := w[:j]
				secondhalf := w[j:]
				result = append(result, firsthalf)
				result = append(result, secondhalf)
			} else {
				result = append(result, w)
			}
		}
	}
	return result
}

func Single_Cote(content string) string {
	new_content := Split_Colon(content, "'")
	if len(new_content) == 0 {
		return content
	}
	colon := false
	str := ""
	for _, arg := range new_content {
		if arg == "'" && !colon {
			colon = true
			continue
		} else if arg == "'" && colon {
			colon = false
			continue
		}

		if colon {
			str += "'" + RemoveSpace(arg) + "'"
		} else {
			str += arg
		}
	}

	return str
}

func RemoveSpace(str string) string {
	return strings.TrimSpace(str)
}

func Split_Colon(s, sep string) []string {
	parts := strings.Split(s, sep)
	result := []string{}
	for i, part := range parts {
		result = append(result, part)
		if i < len(parts)-1 {
			result = append(result, sep) // Add separator back between parts
		}
	}
	return result
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

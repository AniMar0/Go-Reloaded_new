package reload

import (
	"os"
	"os/exec"
	"strings"
	"unicode"
)

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

func RemoveSpace(str string) string {
	return strings.TrimSpace(str)
}

func RemoveSpaceFlag(flag string) string {
	str := ""
	for _, digit := range flag {
		if digit != ' ' {
			str += string(digit)
		}
	}
	return str
}

func SplitLowUpCap(Data string) []string {
	slice := Clean(strings.Split(Data, " "))

	var result []string

	for i := 0; i < len(slice); i++ {
		word := slice[i]
		if i < len(slice)-1 {
			flag := word + " " + slice[i+1]
			if flags.isFlag(flag) {
				result = append(result, flag)
				i++
			} else if i+2 < len(slice)-1 && flags.isFlag(flag+slice[i+2]) {
				result = append(result, flag+slice[i+2])
				i += 2
			} else {
				result = append(result, word)
			}
		}
	}
	if len(slice) >= 2 {
		lastWord := slice[len(slice)-2] + " " + slice[len(slice)-1]

		if flags.isFlag(slice[len(slice)-2] + " " + slice[len(slice)-1]) {
			result = result[:len(result)-1]
			result = append(result, lastWord)
		} else {
			result = append(result, slice[len(slice)-1])
		}
	} else if len(slice) == 1 {
		result = append(result, slice[0])
	}

	return result
}

func Low(word string) string {
	low_word := ""
	for _, char := range word {
		low_word += string(unicode.ToLower(char))
	}
	return low_word
}

func Up(word string) string {
	up_word := ""
	for _, char := range word {
		up_word += string(unicode.ToUpper(char))
	}
	return up_word
}

func Capit(word string) string {
	cap_word := ""
	index := 0
	for i, char := range word {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			if len(word) > index+1 {
				index++
			}
		}
		if i == index {
			cap_word += string(unicode.ToUpper(char))
		} else {
			cap_word += string(unicode.ToLower(char))
		}
	}
	return cap_word
}

func SplitWhiteSpace(s string) []string {
	var arr []string
	var wordStart int
	inWord := false
	for i, char := range s {
		if char == ' ' {
			if inWord {
				arr = append(arr, s[wordStart:i])
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
		arr = append(arr, s[wordStart:])
	}
	return arr
}

func Clean(Data []string) []string {
	var newData []string
	for _, word := range Data {
		if word != "" {
			newData = append(newData, word)
		}
	}
	return newData
}

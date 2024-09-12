package reload

import (
	"os"
	"os/exec"
	"strings"
	"unicode"
)

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

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func SplitUpCapLow(s string) []string {
	var arr []string
	var wordStart int
	inWord := false
	bracket := false
	for i, char := range s {
		if char == '(' {
			bracket = true
		} else if char == ')' {
			bracket = false
		}
		if char == ' ' && !bracket {
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
	for i, char := range word {
		if i == 0 {
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

package reload

import (
	"os"
	"os/exec"
	"strings"
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

func Low(word string) string {
	return strings.ToLower(word)
}

func Up(word string) string {
	return strings.ToUpper(word)
}

func Capit(word string) string {
	cap_word := ""
	if len(word) == 0 || isAlpha(rune(word[0])) {
		// return "", errors.New("You can't capitalize the word: " + word)
		cap_word = strings.ToUpper(string(word[0])) + word[1:]
	}

	return cap_word
}
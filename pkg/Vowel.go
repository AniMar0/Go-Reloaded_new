package reload

import "strings"

func AtoAn(s string) string {
	slice := SplitWhiteSpace(s)
	result := ""

	for i, w := range slice {
		if w == "a" {
			if i != len(slice)-1 {
				nextWord := strings.ToLower(slice[i+1])
				if len(nextWord) > 0 && (nextWord[0] == 'a' || nextWord[0] == 'e' || nextWord[0] == 'i' || nextWord[0] == 'o' || nextWord[0] == 'u') {
					w = "an"
				}
			}
		} else if w == "A" {
			if i != len(slice)-1 {
				nextWord := strings.ToLower(slice[i+1])
				if len(nextWord) > 0 && (nextWord[0] == 'a' || nextWord[0] == 'e' || nextWord[0] == 'i' || nextWord[0] == 'o' || nextWord[0] == 'u') {
					w = "An"
				}
			}
		}
		result += w
		if i < len(slice)-1 {
			result += " "
		}
	}
	return result
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

package reload

import "strings"

func AtoAn(s string) string {
	slice := SplitWhiteSpace(s)
	result := ""

	for i, w := range slice {
		if w == "a" {
			if i != len(slice)-1 {
				nextWord := strings.ToLower(slice[i+1])
				if len(nextWord) > 0 && (nextWord[0] == 'a' || nextWord[0] == 'e' || nextWord[0] == 'i' || nextWord[0] == 'o' || nextWord[0] == 'u' || nextWord[0] == 'h') {
					w = "an"
				}
			}
		} else if w == "A" {
			if i != len(slice)-1 {
				nextWord := strings.ToLower(slice[i+1])
				if len(nextWord) > 0 && (nextWord[0] == 'a' || nextWord[0] == 'e' || nextWord[0] == 'i' || nextWord[0] == 'o' || nextWord[0] == 'u' || nextWord[0] == 'h') {
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

package reload

import "strings"

func AtoAn(s string) string {
	slice := SplitWhiteSpace(s)
	result := ""

	for i, w := range slice {
		if w == "a" || w == "A" {
			if i != len(slice)-1 {
				nextWord := strings.ToLower(slice[i+1])
				if len(nextWord) > 0 && strings.ContainsAny(string(nextWord[0]), "aeiouh") {
					if w == "a" || w == "A" {
						w += "n"
					}
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

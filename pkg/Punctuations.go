package reload

import "strings"

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

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
    var result []string

    slice := SplitWhiteSpace(s)

    for _, word := range slice {
        count := 0
        for i := 0; i < len(word) && strings.ContainsAny(string(word[i]), ".,?!:;"); i++ {
            count++
        }
        if count > 0 {
            FirstWord := word[:count]
            SecondWord := word[count:]
            result = append(result, FirstWord, SecondWord)
        } else {
            result = append(result, word)
        }
    }

    return result
}
package reload

func Single_Cote(content string) string {
	new_content := Split_Single_Cote(content, "'")

	if len(new_content) == 0 {
		return content
	}
	colon := false
	str := ""

	for _, arg := range new_content {
		if arg == "'" && !colon {
			colon = true
		} else if arg == "'" && colon {
			colon = false
		}

		if colon {
			str += RemoveSpace(arg)
		} else {
			str += arg
		}
	}

	return str
}

func Split_Single_Cote(s, sep string) []string {
	s += sep
	Words := []string{}
	last_index := 0
	for i, char := range s {
		if char == '\'' {
			if i == len(s)-1 {
				Words = append(Words, s[last_index:])
			} else if i > 0 && !isAlpha(rune(s[i-1])) && !isAlpha(rune(s[i+1])) {
				Words = append(Words, s[last_index:i])
				Words = append(Words, "'")
				last_index = i + 1

			} else if i == 0 {
				Words = append(Words, string(char))
				last_index = i + 1
			}
		}
	}
	Words = append(Words, s[last_index:])

	return Words
}

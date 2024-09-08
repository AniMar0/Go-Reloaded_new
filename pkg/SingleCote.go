package reload

func Single_Cote(content string) string {
	new_content := Split_Single_Cote(content, "'")

	println(new_content)

	if len(new_content) == 0 {
		return content
	}
	First_colon := false
	last_colon := false
	str := ""

	for _, arg := range new_content {
		if arg == "'" && !First_colon {
			First_colon = true
			continue
		} else if arg == "'" && First_colon {
			First_colon = false
			last_colon = true
			continue
		}

		if last_colon {
			str += "'" + RemoveSpace(arg) + "'"
			last_colon = false
		} else if First_colon {
			str += "'" + arg
		} else {
			str += arg
		}
	}

	return str
}

func Split_Single_Cote(s, sep string) []string {
	s += sep
	Words := []string{}
	add := false
	for i, char := range s {
		if add {
			s = s[i:]
			add = false
		}
		if char == '\'' {
			Words = append(Words, s[:i])
			Words = append(Words, "'")
			add = true
		}
	}
	return Words
}

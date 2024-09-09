package reload

func Single_Cote(content string) string {
	new_content := Split_Single_Cote(content, "'")

	if len(new_content) == 0 {
		return content
	}

	counter := 0
	str := ""
	for i, arg := range new_content {
		
		if arg == "'" {
			counter++
		}

		if counter == 2 {
			str += "'" + RemoveSpace(new_content[i-1]) + "'"
			counter = 0
		} else if counter == 0 {
			str += arg
		} else if i == len(new_content)-1 && counter == 1 {
			str += "'" + arg
		}

	}
	// colon := false
	// str := ""

	// for _, arg := range new_content {
	// 	// println(arg)
	// 	if arg == "'" && !colon {
	// 		colon = true
	// 	} else if arg == "'" && colon {
	// 		colon = false
	// 	}

	// 	if colon {
	// 		str += RemoveSpace(arg)
	// 	} else {
	// 		if arg[len(arg)-1] != ' ' {
	// 			str += arg + " "
	// 		} else {
	// 			str += arg
	// 		}
	// 	}
	// }

	return str
}

func Split_Single_Cote(s, sep string) []string {
	Words := []string{}
	last_index := 0
	str := ""
	for i, char := range s {
		if char == '\'' {
			if i == 0 {
				Words = append(Words, string(char))
				last_index = i + 1
			} else if i > 0 && (!isAlpha(rune(s[i-1])) || !isAlpha(rune(s[i+1]))) {
				str = s[last_index:i]
				Words = append(Words, str)
				Words = append(Words, "'")
				str = ""
				last_index = i + 1

			}
		}
		if i == len(s)-1 {
			str = s[last_index:]
			Words = append(Words, str)
		}
	}

	return Words
}

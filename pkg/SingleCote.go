package reload

import "strings"

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
			continue
		} else if arg == "'" && colon {
			colon = false
			continue
		}

		if colon {
			str += "'" + RemoveSpace(arg) + "'"
		} else {
			str += arg
		}
	}

	return str
}

func Split_Single_Cote(s, sep string) []string {
	parts := strings.Split(s, sep)
	result := []string{}
	for i, part := range parts {
		result = append(result, part)
		if i < len(parts)-1 {
			result = append(result, sep) // Add separator back between parts
		}
	}
	return result
}

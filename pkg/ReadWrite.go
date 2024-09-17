package reload

import (
	"errors"
	"os"
	"strings"
)

func ReadFile() (string, string, error) {
	args := os.Args
	isValid, flag := isValidArgs(args)

	if !isValid {
		return "", "", errors.New("invalid arguments")
	}

	if flag == "--All" {
		Start := Start_Menu(args)
		if !Start {
			return "", "", errors.New(" Exiting program. ")
		}
	}

	fileContent, err := readFileContent(args[1])
	if err != nil {
		return "", "", err
	}

	if fileContent == "" {
		return "", "", errors.New("nothing to read")
	}
	if flag == "--All" {
		return fileContent, "--All", nil
	}
	return fileContent, "", nil
}

func WriteFile(data string) error {
	return os.WriteFile(string(os.Args[2]), []byte(data), 0o644)
}

func readFileContent(fileName string) (string, error) {
	Data, err := os.ReadFile(fileName)
	return string(Data), err
}

func isValidArgs(args []string) (bool, string) {
	if len(args) == 3 && isTxtFile(args[1]) && isTxtFile(args[2]) {
		return true, ""
	} else if len(args) == 4 && isTxtFile(args[1]) && isTxtFile(args[2]) && args[3] == "--All" {
		return true, "--All"
	} else {
		return false, ""
	}
}

func isTxtFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".txt")
}

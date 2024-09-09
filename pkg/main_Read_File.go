package reload

import (
	"bufio"
	"errors"
	"os"
)

func Read_File() (string, error) {
	Args := os.Args
	if !isValidArg(Args) {
		return "", errors.New("invalid Args")
	}

	file, err := os.Open(Args[1])
	if err != nil {
		return "", errors.New("File not Found: " + Args[1])
	}
	defer file.Close()

	// Start the process based on Start_Menu's result
	Start := Start_Menu(Args)
	if !Start {
		return "", errors.New(" Exiting program. ")
	}

	// Using bufio.Scanner to read file line by line
	var fileContent string
	line_vide := 0
	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		if line == "" {
			line_vide++
		}
		fileContent += line + "\n"
	}
	if lines == line_vide {
		return "", errors.New(" Noting to read ")
	} else {
		fileContent = fileContent[:len(fileContent)-1]
	}

	// Checking for scanning errors
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fileContent, nil
}

func isvalidExt(file_name string) bool {
	for i, char := range file_name {
		if char == '.' && i != 0 {
			if file_name[i:] == ".txt" {
				return true
			}
			break
		}
	}
	return false
}

func isValidArg(args []string) bool {
	if len(args) != 3 {
		return false
	}
	if isvalidExt(args[1]) && isvalidExt(args[2]) {
		return true
	}
	return false
}

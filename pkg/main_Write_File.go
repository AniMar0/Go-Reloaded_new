package reload

import (
	"bufio"
	"os"
	"strings"
)

func Write_File(Data string) error {
	file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Split the Data into lines and write them one by one
	lines := strings.Split(Data, "\n") // Assume this splits the content by lines (or use strings.Split if needed)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	// Flush any remaining buffered data to the file
	return writer.Flush()
}

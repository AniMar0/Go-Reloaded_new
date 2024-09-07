package reload

import (
	"fmt"
	"os"
)

func Read_File() string {
	Args := os.Args
	if !isValidArg(Args) {
		fmt.Println("invalid Args")
	}
	file, _ := os.ReadFile(Args[1])

	return string(file)
}

func Modifications(Data string) (string, error) {
	var err error

	Data, err = Convert_To(Data)
	if err != nil {
		return "there is error in Convert_To\n", err
	}

	return Data, err
}

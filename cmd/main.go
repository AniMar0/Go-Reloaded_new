package main

import (
	"fmt"
	"os"

	reload "reload/pkg"
)

func main() {
	Data, flag, err := reload.ReadFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if flag != "--All" {
		Data, err = reload.All_Control(Data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

	} else {
		Data, err = reload.Modifications_Control(Data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	reload.WriteFile(Data)
}

// amine

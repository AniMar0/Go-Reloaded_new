package main

import (
	"fmt"
	"os"

	reload "reload/pkg"
)

func main() {
	
	Data, err := reload.Read_File()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	Data, err = reload.Modifications_Control(Data)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	reload.Write_File(Data)
}

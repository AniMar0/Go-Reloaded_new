package main

import (
	"fmt"

	reload "reload/pkg"
)

func main() {
	Data, err := reload.Read_File()
	if err != nil {
		fmt.Println(err)
		return
	}

	Data, err = reload.Modifications(Data)
	if err != nil {
		fmt.Println(err)
		return
	}

	reload.Write_File(Data)
}

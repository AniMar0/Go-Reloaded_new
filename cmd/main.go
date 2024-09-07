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

	Data = reload.Modifications(Data)

	reload.Write_File(Data)
}

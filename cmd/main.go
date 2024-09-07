package main

import (
	"fmt"
	"strings"

	reload "reload/pkg"
)

func main() {
	Data, err := reload.Read_File()
	if err != nil {
		fmt.Println(err)
		return
	}
	Data_slice := strings.Split(Data, "\n")
	New_Data := ""
	for i := 0; i < len(Data_slice); i++ {
		Data, err = reload.Modifications(Data_slice[i], i+1, len(Data_slice))
		if err != nil {
			fmt.Printf("Error Line %v : %v", i+1, err.Error())
			return
		}
		New_Data += Data + "\n"
	}

	reload.Write_File(New_Data)
}

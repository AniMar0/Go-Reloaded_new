package main

import (
	reload "reload/pkg"
)

func main() {

	Data := reload.Read_File()

	Data = reload.Modifications(Data)

	reload.Write_File(Data)
	
}

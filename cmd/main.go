package main

import reload "reload/pkg"

func main() {
	Data := reload.Read_File()
	Modifications(Data)
	reload.Read_File()
}

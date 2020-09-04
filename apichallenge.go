package main

import (
	"challenge/api/process"

	"github.com/rafael180496/libcore/utility"
)

func main() {
	err := process.Mainprocess()
	if err != nil {
		utility.PrintRed("Error:%s\n", err.Error())
	}
}

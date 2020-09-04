package main

import (
	"challenge/api/process"
	"flag"

	"github.com/rafael180496/libcore/utility"
)

func main() {
	mode := flag.String("mode", "api", "a string")
	path := flag.String("path", "", "a string")
	flag.Parse()
	err := process.Mainprocess(*mode, *path)
	if err != nil {
		utility.PrintRed("Error:%s\n", err.Error())
	}

}

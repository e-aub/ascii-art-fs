package main

import (
	f "ascii-art-fs/functions"
	"os"
)

func main() {
	toWrite, banner := f.ArgsChecker(os.Args)
	if toWrite == "" {
		return
	}
	f.MapFont(banner, f.Minimize(toWrite))
	slicedToWrite := f.Split(toWrite)
	f.Printer(slicedToWrite)

}

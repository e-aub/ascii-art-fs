package functions

import (
	"log"
	"regexp"
)

func ArgsChecker(args []string) (string, string) {
	var banner string
	if len(args) < 2 || len(args) > 3 {
		log.Fatal("Invalid arguments, Usage: go run . [STRING] [BANNER]")
	}
	toWrite := args[1]
	if len(args) == 3 {
		banner = args[2]
		if !regexp.MustCompile(`\.txt$`).Match([]byte(banner)) {
			banner = banner + ".txt"
		}
	} else {
		banner = "standard.txt"
	}
	if banner != "standard.txt" && banner != "shadow.txt" && banner != "thinkertoy.txt" {
		log.Fatalln("invalid banner")
	}
	return toWrite, banner
}

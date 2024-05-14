package functions

import (
	"log"
	"regexp"
)

func ArgsChecker(args []string) (string, string) {
	var banner string
	if len(args) < 2 || len(args) > 3 {
		log.Fatal("\nUsage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
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
	if banner != "standard.txt" && banner != "shadow.txt" && banner != "thinkertoy.txt" && banner != "enigma.txt" {
		log.Fatalln("invalid banner")
	}
	return toWrite, banner
}

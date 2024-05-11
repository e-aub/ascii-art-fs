package functions

import "fmt"

func Printer(toPrint []string) {
	for _, part := range toPrint {
		if part == "\\n" {
			fmt.Println()
			continue
		}
		count := 0
		for count < 8 {
			for _, letter := range part {
				fmt.Print(Font[letter][count])
			}
			fmt.Println()
			count++
		}
	}
}

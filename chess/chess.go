package chess

import (
	"fmt"
)

func Layout() {
	height := readIntParam("Enter the height")
	width := readIntParam("Enter the width")
	symbol := readStringParam("Enter the symbol")

	for i := int16(0); i < height; i++ {
		line := ""
		if (i%2 != 0) {
			line += " "
		}
		for j := int16(0); j < width; j++ {
			line += symbol + " "
		}
		fmt.Println(line)
	}
}

func readIntParam(question string) int16 {
	fmt.Println(question)
	var n int16
    fmt.Scanln(&n)
	return n
}

func readStringParam(question string) string {
	fmt.Println(question)
	var s string
    fmt.Scanln(&s)
	return s
}
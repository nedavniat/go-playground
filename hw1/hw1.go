package main

import (
	"fmt"
	"example.com/chess"
	"example.com/countnums"
	"example.com/card"
	"example.com/fib"
)

func main() {
	gen := fib.FibonacciGen()
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	chess.Layout()
	evenNumbs := countnums.CountNums()
	fmt.Printf("There had been %v positive even numbers", evenNumbs)

	c, err := card.Anonimize()

	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}

}
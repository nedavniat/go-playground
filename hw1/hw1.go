package main

import (
	"fmt"
	"example.com/chess"
	"example.com/countnums"
	"example.com/card"
	"example.com/fib"
)

func main() {
	printMsg("Running task with chess layout")
	chess.Layout()

	printMsg("Running task with count even positive nums")
	evenNumbs := countnums.CountNums()
	fmt.Printf("There had been %v positive even numbers \n", evenNumbs)

	printMsg("Running task with card anonimization")
	c, err := card.Anonimize()

	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}

	printMsg("Generating 6 nums in fib sequence")
	gen := fib.FibonacciGen()
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())

}

func printMsg(msg string) {
	fmt.Println("==========================================")
	fmt.Println(msg)
}
package main

import (
	"fmt"
	"time"
)

//Variadics can be of any type

//Println is our implementation of Println
func Println(args ...interface{}) {
	for _, arg := range args {
		fmt.Printf("%v ", arg)
	}

	fmt.Printf("\n")
}

//Variadics can be of a specific type

//Min returns the lowest value
func Min(args ...int64) int64 {
	var min = int64(^uint(0) >> 1)
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}
	return min
}

//Variadics can even contain functions

//Run contains a list of functions to run
func Run(args ...func()) {
	for _, arg := range args {
		arg()
	}
}

func start() {
	Println("Started...")
}

func end() {
	time.Sleep(time.Second * 2)
	Println("Done")
}

func main() {
	Println("vim-go", 2, &struct{}{})
	Println("Lowest value is", Min(2, 3, 5, 6, 12, 1, 25, 83))
	Run(start, end)
}

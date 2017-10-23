package main

import "fmt"

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

func main() {
	var jazz = &struct{}{}
	Println("vim-go", 2, jazz)
	Println("Lowest value is", Min(2, 3, 5, 6, 12, 1, 25, 83))
}

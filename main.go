package main

import (
	"errors"
	"fmt"
)

func sum(x, y int) (int, error) {
	if x+y > 10 {
		return 0, nil
	}
	return x + y, errors.New("sum is less than 10")
}

func main() {
	x, err := sum(1, 2)
	if err != nil {
		fmt.Println(err)
		panic(err) // stop the program and print the error message
		return
	}
	fmt.Println("Hello, World!", x)
}

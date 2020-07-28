package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func multiply(numbers ...int) int {
	var multiply int = 1
	for _, number := range numbers {
		multiply *= number
	}
	return multiply
}

func main() {
	var numbers []int
	arguments := os.Args[1:]

	for _, argument := range arguments {

		number, err := strconv.Atoi(argument)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Multiply %d\n", multiply(numbers...))
}

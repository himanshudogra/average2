package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...int) float64 {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
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
	fmt.Printf("Average: %f\n", average(numbers...))
}

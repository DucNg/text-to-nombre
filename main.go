package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	numbers "github.com/DucNg/text-to-nombre/lib"
)

func main() {
	numbersSlice := getNumbersSlice()
	frenchNumbers := numbers.ToFrench(numbersSlice)

	fmt.Println(frenchNumbers)
}

func getNumbersSlice() []int {
	if len(os.Args) < 2 {
		fmt.Println("specify array of numbers to convert to french text as cli argument, eg. [0, 1, 5, 10, 11, 15, 20, 21, 30, 35, 50, 51, 68, 70, 75, 99, 100, 101, 105, 111, 123, 168, 171, 175, 199, 200, 201, 555, 999, 1000, 1001, 1111, 1199, 1234, 1999, 2000, 2001, 2020, 2021, 2345, 9999, 10000, 11111, 12345, 123456, 654321, 999999]")
		os.Exit(0)
	}

	input := os.Args[1]

	input = strings.Trim(input, "[]")
	input = strings.ReplaceAll(input, " ", "")
	sliceParam := strings.Split(input, ",")
	var numbers []int

	for i := 0; i < len(sliceParam); i++ {
		number, err := strconv.Atoi(sliceParam[i])
		if err != nil {
			log.Fatalf("error parsing input %v\n", sliceParam[i])
		}

		numbers = append(numbers, number)
	}

	return numbers
}

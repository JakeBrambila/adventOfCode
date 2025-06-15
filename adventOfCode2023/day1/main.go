package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

func main() {
	part1()
}

// brute force double triple loop inefficient answer blah blah blah
func part1() {
	sumSlice := []int{}
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runeLine := []rune(line)

		var tempNum int

		for i := 0; i < len(runeLine); i++ {
			if unicode.IsDigit(runeLine[i]) {
				num := int(runeLine[i] - '0')
				tempNum += num * 10
				break
			}
		}
		for j := len(runeLine) - 1; j >= 0; j-- {
			if unicode.IsDigit(runeLine[j]) {
				num := int(runeLine[j] - '0')
				tempNum += num
				break
			}
		}

		sumSlice = append(sumSlice, tempNum)
	}

	elapsed := time.Since(start)
	fmt.Printf("The calibration value is: %d\n", addAllNums(&sumSlice))
	fmt.Println("Time elapsed: ", elapsed)
}

func addAllNums(numSlice *[]int) int {
	var temp int
	for i := 0; i < len(*numSlice); i++ {
		temp += (*numSlice)[i]
	}

	return temp
}

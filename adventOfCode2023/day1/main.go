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
	fmt.Println()
	part2()
}
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
		sumSlice = append(sumSlice, getNums(line))

	}
	elapsed := time.Since(start)
	fmt.Printf("The calibration value is: %d\n", addAllNums(&sumSlice))
	fmt.Println("Time elapsed: ", elapsed)
}

// brute force double triple loop inefficient answer blah blah blah
func getNums(line string) int {
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

	return tempNum
}

func addAllNums(numSlice *[]int) int {
	var temp int
	for i := 0; i < len(*numSlice); i++ {
		temp += (*numSlice)[i]
	}

	return temp
}

func part2() {
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
		sumSlice = append(sumSlice, getAllNums(line))

	}
	elapsed := time.Since(start)
	fmt.Printf("The calibration value is: %d\n", addAllNums(&sumSlice))
	fmt.Println("Time elapsed: ", elapsed)
}

func getAllNums(line string) int {
	numWords := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	var matches []int

	for i := 0; i < len(line); i++ {
		ch := line[i]

		// Check if it's a digit
		if ch >= '0' && ch <= '9' {
			matches = append(matches, int(ch-'0'))
			continue
		}

		// Check if it matches any number word
		for word, val := range numWords {
			//makes sure index doesn't go out of range and then checks if the next characters match the ones in the map
			//e.g. "fgonexa" starting at index 2 line[2:5] == "one"
			if i+len(word) <= len(line) && line[i:i+len(word)] == word {
				matches = append(matches, val)
				break
			}
		}
	}

	if len(matches) == 0 {
		return 0
	}
	return matches[0]*10 + matches[len(matches)-1]
}

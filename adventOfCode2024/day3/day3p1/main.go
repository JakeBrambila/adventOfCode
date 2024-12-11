package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time since start: ", time.Since(start))
	}()
	byteFile, _ := os.ReadFile("input.txt")

	data := string(byteFile)
	phrase := check(data)
	println(MultiplyNums(phrase))

}

func check(result string) []string {
	allPattern := `(?:mul\(\d+,\d+\)|do\(\)|don't\(\))`
	re, _ := regexp.Compile(allPattern)

	var multiplyPhrases []string
	matches := re.FindAllString(result, -1)

	allowed := true
	for i := 0; i < len(matches); i++ {
		if matches[i] == "don't()" {
			allowed = false
			continue
		}
		if matches[i] == "do()" {
			allowed = true
			continue
		}
		if allowed {
			multiplyPhrases = append(multiplyPhrases, matches[i])
		}
	}
	return multiplyPhrases
}

func MultiplyNums(phrase []string) int {
	var sum int
	pattern := `\d+`
	re, _ := regexp.Compile(pattern)

	for i := 0; i < len(phrase); i++ {
		matches := re.FindAllString(phrase[i], -1)
		num1, _ := strconv.Atoi(matches[0])
		num2, _ := strconv.Atoi(matches[1])
		product := num1 * num2
		sum += product
	}
	return sum

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Program took ", time.Since(start))
	}()
	var safeSum int

	file, err := os.Open("input.txt")
	errorCheck(err, "Error opening file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var numbers []int
		line := scanner.Text()
		substrings := strings.Split(line, " ")
		for _, s := range substrings {
			num, _ := strconv.Atoi(s)
			numbers = append(numbers, num)
		}
		if checkDecreasing(numbers) {
			safeSum++
			continue
		} else if checkIncreasing(numbers) {
			safeSum++
			continue
		} else if checkDecreasingPermutation(numbers) {
			safeSum++
			continue
		} else if checkIncreasingPermutation(numbers) {
			safeSum++
			continue
		}
	}
	println("Safe: ", safeSum)

}

func errorCheck(err error, message string) {
	if err != nil {
		println("Error: ", message)
		return
	}
}

func checkDecreasing(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if i+1 >= len(numbers) {
			break
		}
		if numbers[i]-numbers[i+1] > 0 && numbers[i]-numbers[i+1] < 4 {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkIncreasing(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if i+1 >= len(numbers) {
			break
		}
		if numbers[i+1]-numbers[i] > 0 && numbers[i+1]-numbers[i] < 4 {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkDecreasingPermutation(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		temp := make([]int, len(numbers))
		copy(temp, numbers)
		temp = append(temp[:i], temp[i+1:]...)
		if checkDecreasing(temp) {
			return true
		}
	}
	return false
}

func checkIncreasingPermutation(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		temp := make([]int, len(numbers))
		copy(temp, numbers)
		temp = append(temp[:i], temp[i+1:]...)
		if checkIncreasing(temp) {
			return true
		}
	}
	return false
}

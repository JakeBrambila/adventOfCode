package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	first := []int{}
	second := []int{}
	var sum int

	file, err := os.Open("input.txt")
	ErrorCheck(err, "Error opening file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		word1, err := strconv.Atoi(words[0])
		ErrorCheck(err, "Error converting string to int")
		word2, err := strconv.Atoi(words[1])
		ErrorCheck(err, "Error converting string to int")
		first = append(first, word1)
		second = append(second, word2)
	}
	first = quickSort(first)
	second = quickSort(second)
	for i := 0; i < len(first); i++ {
		sum += int(math.Abs(float64(first[i] - second[i])))
	}
	elapsed := time.Since(start)
	fmt.Printf("Difference is %d\n", sum)
	fmt.Println("Time elapsed: ", elapsed)
}

func ErrorCheck(err error, message string) {
	if err != nil {
		println(message)
		return
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}

	for _, v := range arr[:len(arr)/2] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	for _, v := range arr[len(arr)/2+1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

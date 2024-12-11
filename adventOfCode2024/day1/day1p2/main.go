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
	var sum int
	firstMap := make(map[int]int)
	secondMap := make(map[int]int)

	file, err := os.Open("input.txt")
	ErrorCheck(err, "Error opening file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//read the file line by line and parse the first and second numbers into two separate slices
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		word1, err := strconv.Atoi(words[0])
		ErrorCheck(err, "Error converting string to int")
		word2, err := strconv.Atoi(words[1])
		ErrorCheck(err, "Error converting string to int")
		firstMap[word1]++
		secondMap[word2]++
	}
	for k, v := range firstMap {
		if _, ok := secondMap[k]; ok {
			sum += v * secondMap[k] * k
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Similarity score is %d\n", sum)
	fmt.Println("Time elapsed: ", elapsed)
}
func ErrorCheck(err error, message string) {
	if err != nil {
		println(message)
		return
	}
}

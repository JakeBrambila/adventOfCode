package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const operators string = "+*"

var operation = make(map[int][]int)

func init() {
	file, err := os.Open("input2.txt")
	errorCheck(err, "Error opening file.")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var temp []int
		result := strings.Split(scanner.Text(), ":")

		num, err := strconv.Atoi(result[0])
		errorCheck(err, "Error converting string to int.")

		nums := strings.Split(strings.TrimSpace(result[1]), " ")
		for i := range nums {
			num, err := strconv.Atoi(nums[i])
			errorCheck(err, "Error converting string to int.")

			temp = append(temp, num)
		}
		operation[num] = temp
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed: ", time.Since(start))
	}()
	fmt.Println("Total sum: ", part1())

}

func part1() int {
	sum := 0

	//loop over all keys and values in the map
	for key, value := range operation {
		//gets the size of the combinations 2^(n-1)
		arrSize := len(value) - 1

		//array of all the possible combinations of the + and * operators
		results := generateCombinations(operators, arrSize)

		//loops over the array of combinations
		for i := 0; i < len(results); i++ {
			potentialSum := value[0]
			for j := 0; j < len(results[i]); j++ {
				switch results[i][j] {
				case '+':
					potentialSum += value[j+1]
				case '*':
					potentialSum *= value[j+1]
				}
				if potentialSum > key {
					break
				}
			}

			if potentialSum == key {
				sum += potentialSum
				break
			}
		}
	}
	return sum
}

func generateCombinations(operators string, length int) []string {
	var results []string
	if length <= 0 {
		return results
	}

	var backtrack func(current string)
	backtrack = func(current string) {
		if len(current) == length {
			results = append(results, current)
			return
		}
		for _, op := range operators {
			backtrack(current + string(op))
		}
	}

	backtrack("")
	return results
}

func errorCheck(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

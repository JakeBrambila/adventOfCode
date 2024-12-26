package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var operators = []string{"+", "*", "||"}

var answers []int
var operands [][]int

func init() {
	file, err := os.Open("input2.txt")
	errorCheck(err, "Error opening file.")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var temp []int
		result := strings.Split(scanner.Text(), ":")

		num, err := strconv.Atoi(result[0])
		errorCheck(err, "Error converting string to int.")
		answers = append(answers, num)

		nums := strings.Split(strings.TrimSpace(result[1]), " ")
		for i := range nums {
			num, err := strconv.Atoi(nums[i])
			errorCheck(err, "Error converting string to int.")

			temp = append(temp, num)
		}
		operands = append(operands, temp)
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

	//loop over the length of the list
	for k := 0; k < len(answers); k++ {
		//gets the size of the combinations 2^(n-1)
		arrSize := len(operands[k]) - 1

		//array of all the possible combinations of the + and * operators
		results := generateCombinations(operators, arrSize)

		//loops over the array of combinations
		for i := 0; i < len(results); i++ {
			potentialSum := operands[k][0]
			for j := 0; j < len(results[i]); j++ {
				switch results[i][j] {
				case "+":
					potentialSum += operands[k][j+1]
				case "*":
					potentialSum *= operands[k][j+1]
				case "||":
					//concatenates ex: 97 || 7 = 977
					tempStr := strconv.Itoa(potentialSum)
					tempStr += strconv.Itoa(operands[k][j+1])
					num, err := strconv.Atoi(tempStr)
					errorCheck(err, "Error converting string to int.")
					potentialSum = num
				}
				if potentialSum > answers[k] {
					break
				}
			}

			if potentialSum == answers[k] {
				sum += potentialSum
				break
			}
		}
	}
	return sum
}

func generateCombinations(operators []string, length int) [][]string {
	var results [][]string
	if length <= 0 {
		return results
	}

	var backtrack func(current []string)
	backtrack = func(current []string) {
		if len(current) == length {
			// Make a copy of the current combination to avoid issues with shared references
			combination := make([]string, len(current))
			copy(combination, current)
			results = append(results, combination)
			return
		}
		for _, op := range operators {
			backtrack(append(current, op))
		}
	}

	backtrack([]string{})
	return results
}

func errorCheck(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

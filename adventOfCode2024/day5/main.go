package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var numMap = make(map[int][]int)
var updates = [][]int{}
var unsorted [][]int

func init() {
	file, _ := os.Open("input2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//puts all the numbers into a map in the first half of the input
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		//splits the string by ',' to get the updates all into a 2d int array
		if strings.ContainsRune(scanner.Text(), ',') {
			numsAsStrings := strings.Split(scanner.Text(), ",")
			var nums []int
			for i := 0; i < len(numsAsStrings); i++ {
				num, _ := strconv.Atoi(strings.TrimSpace(numsAsStrings[i]))
				nums = append(nums, num)

			}
			updates = append(updates, nums)
			continue

		}
		nums := strings.Split(scanner.Text(), "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		numMap[num1] = append(numMap[num1], num2)
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed", time.Since(start))
	}()
	part2()
}

func part1() {
	fmt.Println("Sum: ", getUpdateSum())
}

func part2() {
	getUpdateSum()
	for i := 0; i < len(unsorted); i++ {
		line := unsorted[i]
		sort.Slice(line, func(i, j int) bool {
			return isInArray(line[j], numMap[line[i]])
		})
	}
	for i := 0; i < len(unsorted); i++ {
		updates = append(updates, unsorted[i])
	}
	fmt.Println("Sum after sorting: ", getSum())

}

func isInArray(num int, nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			return true
		}
	}
	return false
}

func getSum() int {
	var sum int
	for i := 0; i < len(unsorted); i++ {
		index := len(unsorted[i]) / 2
		sum += unsorted[i][index]
	}
	return sum
}

func getUpdateSum() int {
	var sum int
	for i := 0; i < len(updates); i++ {
		isValid := true
		for j := 0; j < len(updates[i])-1; j++ {
			current := updates[i][j]
			if _, ok := numMap[current]; !ok {
				isValid = false
			}
			for k := j + 1; k < len(updates[i]); k++ {
				if !isInArray(updates[i][k], numMap[current]) {
					isValid = false
				}
			}
			if !isValid {
				unsorted = append(unsorted, updates[i])
				break
			}
		}
		if isValid {
			index := len(updates[i]) / 2
			sum += updates[i][index]
		}
	}
	return sum
}

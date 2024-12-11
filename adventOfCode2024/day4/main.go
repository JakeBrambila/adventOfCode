package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

var crossword []string

func init() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		crossword = append(crossword, scanner.Text())
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed: ", time.Since(start))
	}()
	// part1()
	part2()
}

func part1() {
	var sum int
	sum += findHorizontally(crossword)
	sum += findVertically(crossword)
	sum += findDiagonally(crossword)

	fmt.Println(sum)
}

func part2() {
	fmt.Println("Sum: ", findXmas(crossword))
}

func findHorizontally(line []string) int {
	var sum int
	xmas := regexp.MustCompile("XMAS")
	samx := regexp.MustCompile("SAMX")
	for i := 0; i < len(line); i++ {
		sum += len(xmas.FindAllString(line[i], -1))
		sum += len(samx.FindAllString(line[i], -1))
	}
	return sum
}

func findVertically(line []string) int {
	var sum int
	columns := getColumnsAsStrings(line)
	sum += findHorizontally(columns)

	return sum
}

func getColumnsAsStrings(strings []string) []string {

	// Convert each string into a slice of runes and find the maximum length
	var runeSlices [][]rune
	maxLength := 0
	for _, s := range strings {
		runes := []rune(s)
		runeSlices = append(runeSlices, runes)
		if len(runes) > maxLength {
			maxLength = len(runes)
		}
	}

	// Initialize columns as strings
	columns := make([]string, maxLength)

	// Populate columns
	for _, runes := range runeSlices {
		for colIndex := 0; colIndex < maxLength; colIndex++ {
			if colIndex < len(runes) {
				columns[colIndex] += string(runes[colIndex])
			} else {
				// Handle shorter strings by adding a space
				columns[colIndex] += " "
			}
		}
	}

	return columns
}

func findDiagonally(line []string) int {
	var sum int
	rows := len(line)
	cols := len(line[0])
	//loop that goes through the row
	for i := 0; i < rows-3; i++ {
		for j := 0; j < cols-3; j++ {
			if line[i][j] == 'X' && line[i+1][j+1] == 'M' && line[i+2][j+2] == 'A' && line[i+3][j+3] == 'S' {
				sum++
			}
			if line[i][j] == 'S' && line[i+1][j+1] == 'A' && line[i+2][j+2] == 'M' && line[i+3][j+3] == 'X' {
				sum++
			}
		}
	}
	//loop that goes through the column
	for i := 3; i < rows; i++ {
		for j := 0; j < cols-3; j++ {
			if line[i][j] == 'X' && line[i-1][j+1] == 'M' && line[i-2][j+2] == 'A' && line[i-3][j+3] == 'S' {
				sum++
			}
			if line[i][j] == 'S' && line[i-1][j+1] == 'A' && line[i-2][j+2] == 'M' && line[i-3][j+3] == 'X' {
				sum++
			}
		}
	}
	return sum
}

func findXmas(grid []string) int {
	var sum int
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 'A' {
				if (grid[i+1][j+1] == 'S' && grid[i-1][j-1] == 'M' || grid[i+1][j+1] == 'M' && grid[i-1][j-1] == 'S') && (grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M' || grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S') {
					sum++
				}
			}
		}
	}
	return sum
}

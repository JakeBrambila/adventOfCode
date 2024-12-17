package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Point struct {
	row    int
	column int
}

type Guard struct {
	row       int
	column    int
	direction rune
}

var grid []string
var guard Guard
var visitedLocations = make(map[Point]bool)
var steps int

// aliases for the guard's direction
const (
	Up    = '^'
	Down  = 'V'
	Left  = '<'
	Right = '>'
)

// parses the input to get the guard's location and direction and puts
// the text in a 2d grid which is a []string
func init() {

	file, _ := os.Open("input2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var index int
	for scanner.Scan() {
		if strings.ContainsRune(scanner.Text(), Up) || strings.ContainsRune(scanner.Text(), Left) || strings.ContainsRune(scanner.Text(), Right) || strings.ContainsRune(scanner.Text(), Down) {
			for i := 0; i < len(scanner.Text()); i++ {
				if scanner.Text()[i] == Left {
					guard.direction = Left
					guard.row = index
					guard.column = i
				}
				if scanner.Text()[i] == Right {
					guard.direction = Right
					guard.row = index
					guard.column = i
				}
				if scanner.Text()[i] == Up {
					guard.direction = Up
					guard.row = index
					guard.column = i
				}
				if scanner.Text()[i] == Down {
					guard.direction = Down
					guard.row = index
					guard.column = i
				}
			}
		}
		grid = append(grid, scanner.Text())
		index++
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed", time.Since(start))
	}()
	part2(grid)
}

func part1() {
	for {
		if guard.row == 0 || guard.column == 0 || guard.row == len(grid)-1 || guard.column == len(grid[0])-1 {
			break
		}
		traverseGrid(grid)
	}
	fmt.Println("Visited Locations: ", len(visitedLocations), "Steps", steps)
	steps = 0
}

func part2(gridParam []string) {
	validObstructions := 0
	for tempRow := 0; tempRow < len(gridParam); tempRow++ {
		for tempColumn := 0; tempColumn < len(gridParam[0]); tempColumn++ {
			if gridParam[tempRow][tempColumn] == '.' {
				//create a temporary grid
				tempGrid := make([]string, len(gridParam))
				copy(tempGrid, gridParam)

				//changes the '.' to a '#'
				temp := []rune(tempGrid[tempRow])
				temp[tempColumn] = '#'
				tempGrid[tempRow] = string(temp)

				//resets the guard's position
				guard.row, guard.column, guard.direction = findGuard(grid)
				steps = 0

				//traverse through the grid
				for {
					//checks for an infinite loop
					if steps > len(tempGrid)*len(tempGrid[0]) {
						validObstructions++
						break
					}
					//guard leaves the grid
					if guard.row == 0 || guard.column == 0 || guard.row == len(tempGrid)-1 || guard.column == len(tempGrid[0])-1 {
						break
					}
					traverseGrid(tempGrid)
				}

			}
		}
	}
	fmt.Println("Valid Obstructions: ", validObstructions)
}

// function that moves the guard through the grid, rotates him 90 degrees
// every time he hits a '#' and stops when he leaves the grid
// counts all his visited locations in the visitedLocations map global variable
func traverseGrid(gridParam []string) {
	//starting position to be included
	visitedLocations[Point{guard.row, guard.column}] = true

	if guard.direction == '^' {
		for tempRow := guard.row - 1; tempRow >= 0; tempRow-- {
			if gridParam[tempRow][guard.column] == '#' {
				guard.row = tempRow + 1
				switchDirection(&guard)
				return
			}
			visitedLocations[Point{tempRow, guard.column}] = true
			steps++
		}
		guard.row = 0
	} else if guard.direction == 'V' {
		for tempRow := guard.row + 1; tempRow < len(gridParam); tempRow++ {
			if gridParam[tempRow][guard.column] == '#' {
				guard.row = tempRow - 1
				switchDirection(&guard)
				return
			}
			visitedLocations[Point{tempRow, guard.column}] = true
			steps++
		}
		guard.row = len(gridParam) - 1
	} else if guard.direction == '<' {
		for tempColumn := guard.column - 1; tempColumn >= 0; tempColumn-- {
			if gridParam[guard.row][tempColumn] == '#' {
				guard.column = tempColumn + 1
				switchDirection(&guard)
				return
			}
			visitedLocations[Point{guard.row, tempColumn}] = true
			steps++
		}
		guard.column = 0
	} else if guard.direction == '>' {
		for tempColumn := guard.column + 1; tempColumn < len(gridParam[0]); tempColumn++ {
			if gridParam[guard.row][tempColumn] == '#' {
				guard.column = tempColumn - 1
				switchDirection(&guard)
				return
			}
			visitedLocations[Point{guard.row, tempColumn}] = true
			steps++
		}
		guard.column = len(gridParam[0]) - 1
	}
}

// takes in the reference to a Guard and changes the direction 90 degrees
func switchDirection(guard *Guard) {
	switch (*guard).direction {
	case Up:
		{
			(*guard).direction = Right
		}
	case Right:
		{
			(*guard).direction = Down
		}
	case Down:
		{
			(*guard).direction = Left
		}
	case Left:
		{
			(*guard).direction = Up
		}
	}
}

// gets the original position of the guard
func findGuard(gridParam []string) (int, int, rune) {
	for row := 0; row < len(gridParam); row++ {
		for col := 0; col < len(gridParam[row]); col++ {
			switch gridParam[row][col] {
			case Up, Down, Left, Right:
				return row, col, rune(gridParam[row][col])
			}
		}
	}
	return -1, -1, 0 // Default return if guard not found
}

// function to quickly print the 2d grid
func printGrid(gridParam []string) {
	for i := 0; i < len(gridParam); i++ {
		fmt.Println(gridParam[i])
	}
}

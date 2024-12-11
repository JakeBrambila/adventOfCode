package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row    int
	column int
}

var grid []string
var coordinate Point
var coordinateMap = make(map[Point]bool)
var direction rune
var done bool = false

const (
	Up    = '^'
	Down  = 'V'
	Left  = '<'
	Right = '>'
)

func init() {

	file, _ := os.Open("input2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var index int
	for scanner.Scan() {
		if strings.ContainsRune(scanner.Text(), Up) || strings.ContainsRune(scanner.Text(), Left) || strings.ContainsRune(scanner.Text(), Right) || strings.ContainsRune(scanner.Text(), Down) {
			for i := 0; i < len(scanner.Text()); i++ {
				if scanner.Text()[i] == Left {
					direction = Left
					coordinate.row = index
					coordinate.column = i
				}
				if scanner.Text()[i] == Right {
					direction = Right
					coordinate.row = index
					coordinate.column = i
				}
				if scanner.Text()[i] == Up {
					direction = Up
					coordinate.row = index
					coordinate.column = i
				}
				if scanner.Text()[i] == Down {
					direction = Down
					coordinate.row = index
					coordinate.column = i
				}
			}
		}
		grid = append(grid, scanner.Text())
		index++
	}
}

func main() {
	part1()
}

func part1() {
	fmt.Println("Starting position: ", string(direction))
	fmt.Println("Starting index: ", coordinate.row, ",", coordinate.column)
	//count the starting position
	coordinateMap[coordinate] = true
	for {
		Move()
		if done {
			break
		}
	}
	fmt.Println("Ending places visited: ", len(coordinateMap))
	fmt.Println("Ending index: ", coordinate.row, ",", coordinate.column)
}

//this could be made better with recursion maybe
func Move() {
	if direction == Up {
		for tempRow := coordinate.row - 1; tempRow >= 0; tempRow-- {
			tempPoint := Point{tempRow, coordinate.column}
			if grid[tempRow][coordinate.column] == '#' {
				coordinate.row = tempRow + 1
				switchDirection()
				return
			}
			coordinateMap[tempPoint] = true
		}
		coordinate.row = 0
		done = true
	} else if direction == Right {
		for tempColumn := coordinate.column + 1; tempColumn < len(grid[coordinate.row]); tempColumn++ {
			tempPoint := Point{coordinate.row, tempColumn}
			if grid[coordinate.row][tempColumn] == '#' {
				coordinate.column = tempColumn - 1
				switchDirection()
				return
			}
			coordinateMap[tempPoint] = true
		}
		coordinate.column = len(grid[coordinate.row])
		done = true
	} else if direction == Down {
		for tempRow := coordinate.row; tempRow < len(grid); tempRow++ {
			tempPoint := Point{tempRow, coordinate.column}
			if grid[tempRow][coordinate.column] == '#' {
				coordinate.row = tempRow - 1
				switchDirection()
				return
			}
			coordinateMap[tempPoint] = true
		}
		coordinate.row = len(grid) - 1
		done = true
	} else if direction == Left {
		for tempColumn := coordinate.column - 1; tempColumn >= 0; tempColumn-- {
			tempPoint := Point{coordinate.row, tempColumn}
			if grid[coordinate.row][tempColumn] == '#' {
				coordinate.column = tempColumn + 1
				switchDirection()
				return
			}
			coordinateMap[tempPoint] = true
		}
		coordinate.column = 0
		done = true
	}

}

func switchDirection() {
	switch direction {
	case Up:
		{
			direction = Right
		}
	case Right:
		{
			direction = Down
		}
	case Down:
		{
			direction = Left
		}
	case Left:
		{
			direction = Up
		}
	}
}

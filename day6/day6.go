package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
	"slices"
	"strings"
)

func walk(loc [2]int, dir [2]int) [2]int {
	var newLoc [2]int
	newLoc[0] = loc[0] + dir[0]
	newLoc[1] = loc[1] + dir[1]
	return newLoc
}

func lookForward(matrix []string, loc [2]int,
	dir [2]int) (inBounds, clear bool) {

	nRows := len(matrix) - 1
	nCols := len(matrix[0]) - 1
	nextDir := walk(loc, dir)

	if (nextDir[0] < 0 || nextDir[0] > nRows) ||
		(nextDir[1] < 0 || nextDir[1] > nCols) {
		return false, false
	}

	if matrix[nextDir[0]][nextDir[1]] == '#' {
		return true, false
	}
	return true, true
}

func main() {
	var matrix []string
	var initLoc [2]int
	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	f := helper.OpenFile("./file.txt")

	// first step: create the matrix
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		matrix = append(matrix, scn.Text())
	}

	// second step: look for the guard (^)
	for i, row := range matrix {
		if strings.Contains(row, "^") {
			initLoc[0] = i
			initLoc[1] = strings.Index(row, "^")
		}
	}

	// third step: moving logic
	dir := 0
	loc := initLoc
	path := make(map[int][]int)
	count := 0

	inBounds, clear := lookForward(matrix, loc, directions[dir])
	fmt.Println(inBounds, clear)
	fmt.Println(loc)

	for true {
		for inBounds && clear {
			loc = walk(loc, directions[dir])
			fmt.Println(loc)
			if !slices.Contains(path[loc[0]], loc[1]) {
				path[loc[0]] = append(path[loc[0]], loc[1])
				count++
			}

			inBounds, clear = lookForward(matrix, loc, directions[dir])
			fmt.Printf("after loc: %v: %v %v \n", loc, inBounds, clear)
		}
		// turn
		if !clear {
			if !inBounds {
				break
			}

			if dir == 3 {
				dir = 0
			} else {
				dir++
			}
			inBounds, clear = lookForward(matrix, loc, directions[dir])

			fmt.Printf("next step after change of direction is: %v %v\n",
				inBounds, clear)

		}
	}
	fmt.Println(count)
}

package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
)

const word_day1 = "XMAS"
const word_day2 = "MAS"

func day2(matrix []string) int {
	var count int
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if matrix[i][j] == 'A' {
				if i-1 >= 0 && i+1 < m && j-1 >= 0 && j+1 < n {
					if ((matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S') ||
						(matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M')) &&
						((matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S') ||
							(matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M')) {
						count++
					}

				}
			}
		}
	}
	return count
}
func day1(matrix []string) int {
	var isXMAS bool
	var count int
	m := len(matrix)
	n := len(matrix[0])

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			for _, dir := range directions {
				isXMAS = true
				rowIndex := dir[0]
				colIndex := dir[1]

				for charIndex := 0; charIndex < 4; charIndex++ {
					rowOffset := i + (rowIndex * charIndex)
					colOffset := j + (colIndex * charIndex)

					if rowOffset < 0 || rowOffset > m-1 || colOffset < 0 ||
						colOffset > n-1 {
						isXMAS = false
						break
					}

					if matrix[rowOffset][colOffset] != word_day1[charIndex] {
						isXMAS = false
					}
				}
				if isXMAS {
					count++
				}
			}
		}
	}
	return count
}
func main() {
	f := helper.OpenFile("./file.txt")
	var matrix []string

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		matrix = append(matrix, scn.Text())
	}

	fmt.Printf("day1: %d\n", day1(matrix))
	fmt.Printf("day2: %d\n", day2(matrix))

}

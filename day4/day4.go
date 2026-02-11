package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
)

func main() {
	f := helper.OpenFile("./file.txt")
	var matrix []string

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		matrix = append(matrix, scn.Text())
	}

	fmt.Printf("rows: %d\n", len(matrix))
	fmt.Printf("cols: %d\n", len(matrix[0]))
}

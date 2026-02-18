package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
	"io"
	"slices"
	"strings"
)

func readInput(file io.Reader) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ints := strings.Split(scanner.Text(), "   ")

		num1 := helper.MustAtoi(ints[0])
		num2 := helper.MustAtoi(ints[1])

		left = append(left, num1)
		right = append(right, num2)
	}
	return left, right
}

func main() {
	file := helper.OpenFile("../file.txt")
	defer file.Close()

	left, right := readInput(file)
	var sum int

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		sum += int(helper.Abs(left[i] - right[i]))
	}
	fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(file io.Reader) ([]int, []int, error) {
	left := make([]int, 0)
	right := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ints := strings.Split(scanner.Text(), "   ")

		num1, err := strconv.Atoi(ints[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(ints[1])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, num1)
		right = append(right, num2)
	}
	return left, right, nil
}

func solvePart1(left, right []int) (sum int) {
	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		sum += int(helper.Abs(left[i] - right[i]))
	}
	return sum
}

func solvePart2(left, right []int) (sum int) {
	mapRight := make(map[int]int)

	for _, v := range right {
		mapRight[v] += 1
	}

	for _, v := range left {
		sum += v * mapRight[v]
	}
	return sum
}

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	left, right, err := readInput(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part1: %d", solvePart1(left, right))
	fmt.Printf("Part2: %d", solvePart2(left, right))

}

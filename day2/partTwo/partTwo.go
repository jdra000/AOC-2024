package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/jdra000/AOC-2024/helper"
)

func isSafe(nums []int, retry bool) bool {

	var diff int

	increasing := nums[1] > nums[0]
	for i := 0; i < len(nums)-1; i++ {
		diff = helper.Abs(nums[i] - nums[i+1])
		decLevel := diff < 1 || diff > 3

		disorder := (increasing && nums[i] > nums[i+1]) ||
			(!increasing && nums[i+1] > nums[i])

		// if enters here something is wrong
		if decLevel || disorder {
			if retry { // recursive step
				for i := 0; i <= len(nums)-1; i++ {
					s1 := nums[:i]
					s2 := nums[i+1:]
					if isSafe(slices.Concat(s1, s2), false) {
						return true
					}
				}
				return false
			}
			return false
		}
	}
	return true
}

func main() {
	file := helper.OpenFile("../file.txt")

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var count int
	for scanner.Scan() {
		nums := helper.ToIntArr(strings.Split(scanner.Text(), " "))

		if isSafe(nums, true) {
			count++
		}
	}
	fmt.Println(count)
}

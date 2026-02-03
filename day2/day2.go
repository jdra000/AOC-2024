package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

		if decLevel || disorder {
			if retry {
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
func toInt(str []string) []int {
	var nums []int
	for _, v := range str {
		nums = append(nums, helper.MustAtoi(v))
	}
	return nums
}
func performCalc(file io.Reader) {

	scanner := bufio.NewScanner(file)

	var validCountDay1 int
	var validCountDay2 int
	for scanner.Scan() {
		nums := toInt(strings.Split(scanner.Text(), " "))

		if isSafe(nums, true) {
			validCountDay2++
		}
		if isSafe(nums, false) {
			validCountDay1++
		}
	}
	fmt.Printf("day1: %d - day2: %d\n", validCountDay1, validCountDay2)
}

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	performCalc(file)
}

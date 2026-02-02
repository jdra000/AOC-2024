package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
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
				return isSafe(slices.Delete(nums, i+1, i+2), false)
			}
			return false
		}
	}
	//fmt.Printf("%v: safe \n", nums)
	return true
}
func toInt(str []string) []int {
	var nums []int
	for _, v := range str {
		val, _ := strconv.Atoi(v)
		nums = append(nums, val)
	}
	return nums
}
func readLine(file io.Reader) {

	scanner := bufio.NewScanner(file)

	// solving part 1
	var validCount int
	for scanner.Scan() {
		nums := toInt(strings.Split(scanner.Text(), " "))

		if isSafe(nums, true) {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	readLine(file)
}

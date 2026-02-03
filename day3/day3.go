package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/jdra000/AOC-2024/helper"
)

func toInt(str []string) []int {
	var nums []int
	for _, v := range str {
		nums = append(nums, helper.MustAtoi(v))
	}
	return nums
}

func performCalc(file io.Reader) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	var count int
	var validStrings []string

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		validStrings = re.FindAllString(line, -1)

		for _, v := range validStrings {
			re2 := regexp.MustCompile(`\d+`)
			validNumbers := toInt(re2.FindAllString(v, -1))
			count += validNumbers[0] * validNumbers[1]
		}
	}
	fmt.Println(count)
}

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	performCalc(file)
}

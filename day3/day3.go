package main

import (
	"bufio"
	"fmt"
    "regexp"
    "strings"

	"github.com/jdra000/AOC-2024/helper"
)

func toInt(str []string) (nums []int) {
	for _, v := range str {
		nums = append(nums, helper.MustAtoi(v))
	}
	return nums
}

func calc(str string, re *regexp.Regexp) (count int) {
    number_regex := regexp.MustCompile(`\d+`)
    validMuls := re.FindAllString(str, -1)

    for _, mul := range validMuls {
        validNumbers := toInt(number_regex.FindAllString(mul, -1))
        count += validNumbers[0] * validNumbers[1]
    }

    return count
}

func partOne(f string, re *regexp.Regexp) (count int) {

    file := helper.OpenFile(f)
    defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
        count += calc(line, re)
	}
	return count
}

func partTwo(f string, re *regexp.Regexp) (count int) {

    file := helper.OpenFile(f)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    do_regex := regexp.MustCompile(`do\(\)`)

    for scanner.Scan(){
        line := scanner.Text()

        split_str := strings.Split(line, "don't()")

        for i, str := range split_str {
            if i == 0 {
                count += calc(str, re)
		    }

            if strings.Contains(str, "do()"){
                index := do_regex.FindStringIndex(str)
                count += calc(str[index[1]:], re)
		    }
        }
    }
    return count
}

func main() {
    f := "./file.txt"
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)

	fmt.Printf("results day1: %d\n", partOne(f, regex))
	fmt.Printf("results day2: %d\n", partTwo(f, regex))
}

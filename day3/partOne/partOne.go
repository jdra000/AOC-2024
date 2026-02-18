package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
	"regexp"
)

func main() {
	var count int
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	numRegex := regexp.MustCompile(`\d+`)

	f := helper.OpenFile("../file.txt")
	defer f.Close()
	scn := bufio.NewScanner(f)

	for scn.Scan() {
		validMuls := regex.FindAllString(scn.Text(), -1)

		for _, mul := range validMuls {
			validNums := helper.ToIntArr(numRegex.FindAllString(mul, -1))
			count += validNums[0] * validNums[1]
		}
	}
	fmt.Println(count)
}

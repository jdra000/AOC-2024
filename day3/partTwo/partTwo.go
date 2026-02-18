package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
	"regexp"
	"strings"
)

var mulRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)
var numRegex = regexp.MustCompile(`\d+`)

func calc(section string) int {
	var count int
	validMuls := mulRegex.FindAllString(section, -1)

	for _, mul := range validMuls {
		validNums := helper.ToIntArr(numRegex.FindAllString(mul, -1))
		count += validNums[0] * validNums[1]
	}
	return count
}

func main() {
	var count int
	doRegex := regexp.MustCompile(`do\(\)`)
	enabled := true

	f := helper.OpenFile("../file.txt")
	defer f.Close()
	scn := bufio.NewScanner(f)

	for scn.Scan() {
		criticalSections := strings.Split(scn.Text(), "don't()")

		for i, section := range criticalSections {

			/* check at the end of the line if we can operate or not on the
			first section of the next line */
			if i == len(criticalSections)-1 {
				if !strings.Contains(section, "do()") {
					enabled = false
				} else {
					enabled = true
				}
			}

			validSection := section

			if i == 0 {
				if enabled {
					count += calc(validSection)
				} else if strings.Contains(section, "do()") {

					index := doRegex.FindStringIndex(section)
					validSection = section[index[1]:]
					count += calc(validSection)
				}
			}

			if i != 0 && strings.Contains(section, "do()") {
				index := doRegex.FindStringIndex(section)
				validSection = section[index[1]:]
				count += calc(validSection)
			}
		}
	}
	fmt.Println(count)
}

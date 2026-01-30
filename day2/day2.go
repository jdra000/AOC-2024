package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jdra000/AOC-2024/helper"
)

func readLine(file io.Reader) {

	scanner := bufio.NewScanner(file)

	// solving part 1
	var validCount int
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")

		// logic
		var diff int
		inc, dec := false, false
		flag := true

		for i, v := range nums {
			value, _ := strconv.Atoi(v)
			nextValue, _ := strconv.Atoi(nums[i+1])

			diff = helper.Abs(value - nextValue)
			if diff < 1 || diff > 3 {
				flag = false
			}

			if inc || dec {
				if inc && value > nextValue {
					flag = false
				} else if dec && nextValue > value {
					flag = false
				}
			}

			if value < nextValue {
				inc = true
			} else {
				dec = true
			}

			if (i + 2) == len(nums)-1 {
				break
			}
		}
		if flag {
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

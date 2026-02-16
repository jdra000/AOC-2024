package main

import (
	"bufio"
	"fmt"
	"github.com/jdra000/AOC-2024/helper"
	"slices"
	"strings"
)

func arrInt(arr []string) []int {
	var intArr []int
	for _, v := range arr {
		intArr = append(intArr, helper.MustAtoi(v))
	}
	return intArr
}

func compareIntArr(arr1, arr2 []int) bool {
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}
	return true
}

func main() {
	f := helper.OpenFile("./file.txt")
	rules := make(map[int][]int)
	var rule string

	scan := bufio.NewScanner(f)

	// handle rules
	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		rule = scan.Text()
		str := strings.Split(rule, "|")
		nums := arrInt(str)

		rules[nums[0]] = append(rules[nums[0]], nums[1])
	}

	// handle updates
	var count, middle int

	for scan.Scan() {

		str := strings.Split(scan.Text(), ",")
		nums := arrInt(str)

		numscpy := slices.Clone(nums)
		numsDoublecpy := slices.Clone(nums)

		for true {
			inOrder := true
			predecessors := []int{}
			numscpy = slices.Clone(numsDoublecpy)

			for i, v := range numscpy {

				for _, p := range predecessors {
					index := slices.Index(numscpy, p)
					if slices.Contains(rules[v], p) {
						numsDoublecpy[i] = p
						numsDoublecpy[index] = v
						inOrder = false

						break
					}
				}
				if !inOrder {
					break
				}
				predecessors = append(predecessors, v)
			}

			if inOrder {
				break
			}
		}

		fmt.Printf("before: %v, after: %v\n", nums, numsDoublecpy)

		if !compareIntArr(numsDoublecpy, nums) {
			middle = numsDoublecpy[(len(numsDoublecpy)-1)/2]
			fmt.Println(middle)
			count += middle
		}
	}

	fmt.Printf("result: %d\n", count)
}

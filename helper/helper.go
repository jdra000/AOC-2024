package helper

import (
	"log"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MustAtoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

package helper

import (
	"log"
	"strconv"
    "os"
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

func OpenFile(s string) (f *os.File){
   f, err := os.Open(s)
   if err != nil {
        log.Fatal(err)
   }
   return f
}

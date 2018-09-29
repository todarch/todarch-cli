package util

import (
	"fmt"
	"os"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func SayLastWords(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

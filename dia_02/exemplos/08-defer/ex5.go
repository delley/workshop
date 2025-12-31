package main

import (
	"fmt"
)

func contador() (n int) {
	defer func() {
		n++
	}()
	return 10
}

func main() {
	fmt.Println(contador())
}

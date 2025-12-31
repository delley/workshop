package main

import "fmt"

func lifo() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	lifo()
}

package main

import "fmt"

func soma(a, b int) int {
	defer fmt.Println("fim da função soma")
	return a + b
}

func main() {
	fmt.Println(soma(3, 4))
}

package main

import "fmt"

func main() {
	var a, b = 2, 4
	c := (a + b)
	if c == 6 {
		fmt.Println("Sua soma está correta.")
		return
	}
	//uma forma de fazer se não
	fmt.Println("Sua soma está errada.")
}

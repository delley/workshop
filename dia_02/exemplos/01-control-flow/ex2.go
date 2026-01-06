package main

import "fmt"

func main() {
	if 11%2 == 0 {
		fmt.Println("É par.")
	} else {
		fmt.Println("É impar.")
	}

	if num := 11; num < 0 {
		fmt.Println(num, "É negativo.")
	} else if num < 10 {
		fmt.Println(num, "Tem um dígito.")
	} else if num > 10 {
		fmt.Println(num, "É maior que 10.")
	} else {
		fmt.Println(num, "Tem vários dígitos.")
	}
}

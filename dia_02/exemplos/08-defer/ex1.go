package main

import "fmt"

func exemplo() {
	defer fmt.Println("mundo!")
	fmt.Print("Olá ")
}

func main() {
	exemplo()
}

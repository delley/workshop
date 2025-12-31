package main

import (
	"fmt"
	"os"
)

func OpenFile() error {
	file, err := os.Open("dados.txt")
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
		fmt.Println("Arquivo fechado com sucesso!")
	}()

	// Simulando operações com o arquivo
	fmt.Println("Arquivo aberto com sucesso:", file.Name())
	return nil
}

func main() {
	OpenFile()
}

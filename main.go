package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var textForFiles string

	readInput(&textForFiles)

	fmt.Println(textForFiles)
}

func readInput(textToWrite *string) bool {
	fmt.Println("Insira o texto que deseja nos arquivos:")

	reader := bufio.NewReader(os.Stdin)
	textForFiles, error := reader.ReadString('\n')

	*textToWrite = textForFiles

	if error != nil {
		log.Fatal(error)

		return false
	}

	return true
}

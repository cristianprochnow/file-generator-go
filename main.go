package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var textForFiles string
	var filesAmount int

	readInput(&textForFiles)
	readFilesAmount(&filesAmount)

	fmt.Println(textForFiles, filesAmount)
}

func readFilesAmount(filesAmount *int) {
	fmt.Println("Insira a quantidade de arquivos:")
	fmt.Scanf("%d", &*filesAmount)
}

func readInput(textToWrite *string) {
	fmt.Println("Insira o texto que deseja nos arquivos:")

	reader := bufio.NewReader(os.Stdin)
	textForFiles, error := reader.ReadString('\n')

	*textToWrite = textForFiles

	if error != nil {
		log.Fatal(error)
	}
}

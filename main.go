package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var textForFiles, checkOutputFlag, fileNameFlag string
	var filesAmount int

	space()
	readInput(&textForFiles)
	space()
	readFilesAmount(&filesAmount)
	space()
	checkToShowOutput(&checkOutputFlag)
	space()
	readFileName(&fileNameFlag)
	space()

	fmt.Println(textForFiles, filesAmount, checkOutputFlag, fileNameFlag)
}

func checkToShowOutput(checkOutput *string) {
	fmt.Println("Mostrar o texto total na tela? [s - sim / n - não]")

	*checkOutput = strings.ToLower(read())
}

func space() {
	fmt.Println("")
}

func readFilesAmount(filesAmount *int) {
	fmt.Println("Insira a quantidade de arquivos:")
	fmt.Scanf("%d", &*filesAmount)
}

func readFileName(fileNameFlag *string) {
	fmt.Println("Qual das opções deseja em relação ao nome do arquivo:")
	fmt.Println("[d] - Data atual")
	fmt.Println("[p] - Primeira palavra do texto")
	fmt.Println("[l] - Primeira letra do texto")

	*fileNameFlag = strings.ToLower(read())
}

func readInput(textToWrite *string) {
	fmt.Println("Insira o texto que deseja nos arquivos:")

	*textToWrite = read()
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	textForFiles, error := reader.ReadString('\n')

	if error != nil {
		log.Fatal(error)

		return ""
	}

	return textForFiles
}

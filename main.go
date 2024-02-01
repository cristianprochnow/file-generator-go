package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

	buildFiles(fileNameFlag, textForFiles, filesAmount)
	printContent(textForFiles, checkOutputFlag)
}

func getFilePath() string {
	return "./files/"
}

func printContent(content string, outputFlag string) {
	if outputFlag == "s" {
		space()
		fmt.Println("Conteúdo inserido: ", content)
		space()
	}
}

func getFileName(chosedFileOption string, writtenText string) string {
	var fileName string

	switch chosedFileOption {
	case "p":
		var words []string

		for index, word := range writtenText {
			words[index] = string(word)

			if word == ' ' {
				break
			}
		}

		fileName = strings.Join(words, "")
		break

	case "l":
		fileName = string(writtenText[0])
		break

	case "d":
	default:
		fileName = time.Now().Format("2006-01-02-15-04-05")
		break
	}

	return fileName
}

func buildFiles(chosedFileOption string, writtenText string, filesAmount int) {
	filePath := getFilePath() + getFileName(chosedFileOption, writtenText)
	byteText := []byte(writtenText)

	var index int
	for index = 0; index < filesAmount; index++ {
		fullFileName := filePath + "_(" + strconv.Itoa(index) + ").txt"

		if !fileExists(fullFileName) {
			initFile(fullFileName)
		}

		writeFile(fullFileName, byteText)
	}
}

func writeFile(filePath string, content []byte) {
	message := os.WriteFile(filePath, content, 0644)

	check(message)
}

func initFile(fullFileName string) {
	fileObject, message := os.Create(fullFileName)

	check(message)

	defer fileObject.Close()
}

func check(message error) {
	if message != nil {
		panic(message)
	}
}

func fileExists(filePath string) bool {
	_, errorMessage := os.Stat(filePath)

	return !errors.Is(errorMessage, os.ErrNotExist)
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
	fmt.Println("[d] - Data atual (Padrão)")
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

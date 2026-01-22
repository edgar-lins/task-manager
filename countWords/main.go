package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Section 1
	fileData, err := os.ReadFile("countWords/text.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	palavras := strings.Fields(string(fileData))

	count := []string{}
	for _, palavra := range palavras {
		count = append(count, palavra)
	}

	fmt.Printf("Palavras no ficheiro: %v\n", palavras)
	fmt.Printf("Número de palavras no ficheiro: %d\n", len(count))
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	reader, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("Erro ao ler o diretório:", err)
		return
	}

	txt := []string{".txt"}
	csv := []string{".csv"}
	json := []string{".json"}
	fmt.Println(reader)
	for _, file := range reader {
		// Verifica se o ficheiro termina com algum dos sufixos
		if strings.Contains(strings.Join(txt, ","), file.Name()[len(file.Name())-3:]) { // -3: para pegar os últimos 3 caracteres
			fmt.Println("Ficheiro:", file.Name())
			// Mover para pasta "files"
			err := os.Rename(file.Name(), "files/txt/"+file.Name())
			if err != nil {
				fmt.Println("Erro ao mover o ficheiro:", err)
			} else {
				fmt.Println("Ficheiro movido com sucesso!")
			}
		}

		if strings.Contains(strings.Join(csv, ","), file.Name()[len(file.Name())-3:]) { // -3: para pegar os últimos 3 caracteres
			fmt.Println("Ficheiro:", file.Name())
			// Mover para pasta "files"
			err := os.Rename(file.Name(), "files/csv/"+file.Name())
			if err != nil {
				fmt.Println("Erro ao mover o ficheiro:", err)
			} else {
				fmt.Println("Ficheiro movido com sucesso!")
			}
		}

		if strings.Contains(strings.Join(json, ","), file.Name()[len(file.Name())-3:]) { // -3: para pegar os últimos 3 caracteres
			fmt.Println("Ficheiro:", file.Name())
			// Mover para pasta "files"
			err := os.Rename(file.Name(), "files/json/"+file.Name())
			if err != nil {
				fmt.Println("Erro ao mover o ficheiro:", err)
			} else {
				fmt.Println("Ficheiro movido com sucesso!")
			}
		}
	}
}

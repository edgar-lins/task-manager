package main

import "fmt"

// 1. A STRUCT: Aqui definimos a "forma" da nossa Tarefa.
// Diferete de JS, precisamos dizer o tipo de cada campo.
type Task struct {
	ID          int
	Description string
	Done        bool
}

func main() {
	fmt.Println("Bem-vindo ao Task Manager em go! 🚀")

	// 2. Criando uma tarefa de exemplo
	// Usamos := para o Go advinhar que isso é uma variável nova
	myTask := Task{
		ID:          1,
		Description: "Aprender a sintaxe básica de Go",
		Done:        false,
	}

	// 3. Printando a tarefa
	// %+v é um truque do Go para mostrar o nome dos campos e os valores
	fmt.Printf("Minha tarefa: %+v\n", myTask)

	// 4. O SLICE: Criando uma lista de tarefas
	// []Task significa "uma lista onde só entram coisas do tipo Task"
	tasks := []Task{
		myTask,
		{ID: 2, Description: "Configurar o VS COde", Done: true},
		{ID: 3, Description: "Dominar structs em Go", Done: false},
	}

	fmt.Println("\n Lista de Tarefas:")
	// 5. O LOOP: Iterando sobre a lista (o "for" do Go)
	// range retorna o índice e o valor. Usamos _ para ignorar o índice
	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("[%s] %s\n", status, t.Description)
	}

}

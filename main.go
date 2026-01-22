package main

import (
	"encoding/json"
	"os"
)

// 1. A STRUCT: Aqui definimos a "forma" da nossa Tarefa.
// Diferete de JS, precisamos dizer o tipo de cada campo.
type Task struct {
	ID          int
	Description string
	Done        bool
}

// func main() {
// 	// Nosso "banco de dados" em memória
// 	tasks := loadTasks()

// 	// Preparando o leitor de input (Lê da entrada padrão do sistema - teclado)
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("🚀 Task Manager CLI")
// 	fmt.Println("---------------------")

// 	// LOOP INFINITO: O programa só para quando chamarmos 'break' ou 'return'
// 	for {
// 		// 1. Mostrar o menu
// 		fmt.Println("\nEscolha uma opção:")
// 		fmt.Println("1. Listar Tarefas")
// 		fmt.Println("2. Adicionar Tarefa")
// 		fmt.Println("3. Sair")
// 		fmt.Print("Digite o número: ")

// 		// 2. Ler a opção do usuário
// 		// Lê até encontrar uma quebra de linha ('\n')
// 		input, _ := reader.ReadString('\n')
// 		// Limpa espaços e o 'enter' do final para não dar erro na comparação
// 		input = strings.TrimSpace(input)

// 		// 3. Processar a escolha (Switch Case)
// 		switch input {
// 		case "1":
// 			fmt.Println("\n--- Suas Tarefas ---")
// 			// Se a lista estiver vazia
// 			if len(tasks) == 0 {
// 				fmt.Println("Nenhuma tarefa encontrada.")
// 			}
// 			//Mostra as tarefas
// 			for _, t := range tasks {
// 				status := "[ ]"
// 				if t.Done {
// 					status = "[x]"
// 				}
// 				fmt.Printf("%d. %s %s\n", t.ID, status, t.Description)
// 			}

// 		case "2":
// 			fmt.Print("Digite a descrição da nova tarefa: ")
// 			// Lê a defscrição da tarefa
// 			text, _ := reader.ReadString('\n')
// 			text = strings.TrimSpace(text)

// 			// --- 🛑 SEU CÓDIGO AQUI 🛑 ---
// 			// Missão:
// 			// 1. Crie uma nova variável do tipo Task.
// 			// 2. O ID pode ser o tamanho da lista +1 (len(tasks) + 1).
// 			// 3. Adicione (append) essa tarefa na lista 'tasks'.
// 			newTask := Task{
// 				ID:          len(tasks) + 1,
// 				Description: text,
// 				Done:        false,
// 			}

// 			tasks = append(tasks, newTask)

// 			// Dica: Lembre-se de como fizemos no código anterior:
// 			// novaTarefa := Task{...}
// 			// tasks = append(tasks, novaTarefa)
// 			saveTasks(tasks)

// 			fmt.Println("Tarefa adicionada com sucesso!")

// 		case "3":
// 			fmt.Println("Saindo... Até mais!")
// 			return // Encerra a função main e o programa

// 		default:
// 			fmt.Println("Opção inválida, tente novamente.")
// 		}
// 	}
// }

// Função para guardar as tarefas no ficheiro "tasks.json"
func saveTasks(tasks []Task) {
	// MarshalIndent transforma a struct em texto JSON bonitinho (com identação)
	data, _ := json.MarshalIndent(tasks, "", "  ")
	// 0644 é a permissão de leitura/escrita do ficheiro
	os.WriteFile("tasks.json", data, 0644)
}

// Função para carregar as tarefas quando o programa inicia
func loadTasks() []Task {
	// Tenta ler o ficheiro
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		// Se der erro (ex: ficheiro não existe), devolve lista vazia
		return []Task{}
	}
	var loadedTasks []Task
	// Unmarshal transforma o texto JSON de voltar em struct Go
	json.Unmarshal(data, &loadedTasks)
	return loadedTasks
}

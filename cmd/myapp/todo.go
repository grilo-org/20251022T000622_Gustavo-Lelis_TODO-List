package main

import (
	"TODO/internal/comandos"
	"TODO/internal/repositorio"
	"TODO/internal/task_do"
	"log"
)

func main() {
	todos := task_do.Todo{}

	storage := repositorio.NewStorage[task_do.Todo]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		log.Fatalf("Warning: Nao foi carregados os dados.")
	}

	cmdFlags := comandos.NewCmdFlags()
	cmdFlags.Execute(&todos)

	err = storage.Save(todos)
	if err != nil {
		log.Fatalf("Error ao salvar os dados: %v\n", err)
	}
}

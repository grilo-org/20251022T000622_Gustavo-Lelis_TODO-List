package task_do

import (
	"testing"
)

func dadosSalvos(t *testing.T, esperado string, task Todo) {
	t.Helper()

	for _, dados := range task {
		if dados.Descricao != esperado {
			t.Errorf("esperado: %s, resposta: %s", esperado, dados.Descricao)
		}
	}
}

func TestAddTask(t *testing.T) {

	t.Run("adicionar task", func(t *testing.T) {
		task := Todo{}
		task.AddTask("Treino Amanha")

		esperado := "Treino Amanha"

		dadosSalvos(t, esperado, task)
	})

	t.Run("adicionar task em maiusculo", func(t *testing.T) {
		task := Todo{}
		task.AddTask("JOGAR O JOGO CS")

		esperado := "JOGAR O JOGO CS"

		dadosSalvos(t, esperado, task)
	})
}

func dadosRemovidos(t *testing.T, tamanho int, task Todo) {
	t.Helper()

	if tamanho-1 != len(task) {
		t.Errorf("Tamanho Original: %d, Tamanho Atual: %d", tamanho, len(task))
	}
}

func dadosVazios(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("e esperado que um erro seja obtido")
	}
}

func TestRemoveTask(t *testing.T) {
	t.Run("remover task", func(t *testing.T) {
		task := Todo{}
		task.AddTask("Treino Amanha")
		task.AddTask("JOGAR O JOGO CS")

		tamanho := len(task)

		err := task.RemoveTask(0)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		dadosRemovidos(t, tamanho, task)
	})

	t.Run("todo vazia", func(t *testing.T) {
		task := Todo{}

		err := task.RemoveTask(0)

		dadosVazios(t, err)

	})

}

func taskAtualizada(t *testing.T, id int, esperado string, task Todo) {
	t.Helper()

	if task[id].Descricao != esperado {
		t.Fatalf("esperado: %s, %s", esperado, task[id].Descricao)
	}
}

func TestUpdateTask(t *testing.T) {
	t.Run("atualizar task", func(t *testing.T) {
		task := Todo{}

		task.AddTask("Treino Amanha")
		task.AddTask("JOGAR O JOGO CS")
		id := 0
		err := task.UpdateTask(id, "Treinar O JOGO CS")
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		esperado := "Treinar O JOGO CS"

		taskAtualizada(t, id, esperado, task)
	})

	t.Run("atualizar task vazia", func(t *testing.T) {
		task := Todo{}

		err := task.UpdateTask(1, "Treinar O JOGO CS")

		dadosVazios(t, err)

	})

	t.Run("atualizar task fora do indice", func(t *testing.T) {
		task := Todo{}

		task.AddTask("Treino Amanha")
		task.AddTask("JOGAR O JOGO CS")
		id := 5
		err := task.UpdateTask(id, "JOGAR CS")

		dadosVazios(t, err)

	})
}

func statusAtualizado(t *testing.T, id int, status bool, task Todo) {
	t.Helper()

	if task[id].Status != status {
		t.Errorf("resultado: %v, esperado: %v", task[id].Status, status)
	}
}

func TestTodo_UpdateStatus(t *testing.T) {
	t.Run("atualizar status da task", func(t *testing.T) {
		task := Todo{}

		task.AddTask("Treino Amanha")
		id := 0
		err := task.UpdateStatus(id)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		statusAtualizado(t, id, true, task)
	})

	t.Run("atualizar status da task vazia", func(t *testing.T) {
		task := Todo{}

		id := 0
		err := task.UpdateStatus(id)
		dadosVazios(t, err)
	})
}

package task_do

import (
	"fmt"
	"github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

const (
	ErrVazio = ErrTodo("Invalid index")
)

func (e ErrTodo) Error() string {
	return string(e)
}

type ErrTodo string

type Task struct {
	Descricao      string     `json:"descricao"`
	Status         bool       `json:"bool"`
	HorarioInicial time.Time  `json:"horarioInicial"`
	HorarioFinal   *time.Time `json:"horarioFinal"`
}

type Todo []Task // -> Lista de tasks

// -> Recebo a descricao da task e adiciono a lista
func (t *Todo) AddTask(descricao string) {

	task := Task{
		Descricao:      descricao,
		Status:         false,
		HorarioInicial: time.Now(),
		HorarioFinal:   nil,
	}

	*t = append(*t, task)
}

func (t *Todo) validaIndex(id int) error {

	if id < 0 || id >= len(*t) {
		err := ErrVazio
		fmt.Println(err)
		return err
	}
	return nil
}

func (t *Todo) RemoveTask(id int) error {
	// -> Cria uma cópia da lista atual para manipular antes de sobrescrever a original.
	todos := *t

	if err := t.validaIndex(id); err != nil {
		return err
	}

	// -> Remove a task no índice especificado
	*t = append(todos[:id], todos[id+1:]...)

	return nil
}

func (t *Todo) UpdateTask(id int, descricao string) error {
	if err := t.validaIndex(id); err != nil {
		return err
	}

	(*t)[id].Descricao = descricao

	return nil
}

func (t *Todo) UpdateStatus(id int) error {
	// -> Cria uma cópia da lista atual para manipular antes de sobrescrever a original.
	todos := *t
	if err := t.validaIndex(id); err != nil {
		return err
	}

	isCompleted := todos[id].Status

	if !isCompleted {
		taskCompleta := time.Now()
		todos[id].HorarioFinal = &taskCompleta
	}

	todos[id].Status = !isCompleted

	return nil
}

func (t *Todo) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Descricao", "Status", "Horario Inicial", "Horario Final")

	for index, t := range *t {
		status := "❌"
		horarioFinal := ""

		if t.Status {
			status = "✅"
			if t.HorarioFinal != nil {
				horarioFinal = t.HorarioFinal.Format(time.RFC1123)

			}

		}

		table.AddRow(strconv.Itoa(index), t.Descricao, status, t.HorarioInicial.Format(time.RFC1123), horarioFinal)

	}

	table.Render()
}

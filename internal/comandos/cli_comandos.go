package comandos

import (
	"TODO/internal/task_do"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add  string
	Del  int
	Edit string
	Stts int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Adicionar uma nova task")
	flag.StringVar(&cf.Edit, "edit", "", "Para editar especifique o id e a descricao. id:descricao")
	flag.IntVar(&cf.Del, "del", -1, "Para deletar especifique o id")
	flag.IntVar(&cf.Stts, "stts", -1, "Para atualizar o status especifique o id")
	flag.BoolVar(&cf.List, "list", false, "Listar todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *task_do.Todo) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.AddTask(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		err = todos.UpdateTask(index, parts[1])
		if err != nil {
			fmt.Println("Error:", err)
		}

	case cf.Stts != -1:
		err := todos.UpdateStatus(cf.Stts)
		if err != nil {
			fmt.Println("Error:", err)
		}

	case cf.Del != -1:
		err := todos.RemoveTask(cf.Del)
		if err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Invalid command")
	}
}

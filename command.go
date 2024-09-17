package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Deadline string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Sort string
	Ascend bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Deadline, "deadline", "", "Specify deadline for the new todo (e.g., '24h', '1d')")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.StringVar(&cf.Sort, "sort", "", "Sort todos by 'title' or 'deadline" )
	flag.BoolVar(&cf.Ascend, "ascend", true, "Sort by ascending or descending")

	flag.Parse()

	return &cf
}

// Execute method in CmdFlags
func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		if cf.Sort != "" {
			err := todos.sort(cf.Sort, cf.Ascend)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		todos.print()
	case cf.Add != "":
		duration := ""
		if cf.Deadline != "" {
			duration = cf.Deadline
		}
		todos.add(cf.Add, duration)
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

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	case cf.Sort != "":
		err := todos.sort(cf.Sort, cf.Ascend)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		todos.print() // Print sorted list

	default:
		fmt.Println("Invalid command")
	}
}



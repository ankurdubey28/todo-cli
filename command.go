package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo item")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo item by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo item by index:description")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle completion status of a todo item by index")
	flag.BoolVar(&cf.List, "list", false, "List all todo items") // Fixed flag name to 'list'

	flag.Parse()
	return &cf
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid input format for edit. Use index:description")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		if err := todos.toggle(cf.Toggle); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case cf.Delete != -1:
		if err := todos.delete(cf.Delete); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Error: No valid command provided")
		os.Exit(1)
	}
}

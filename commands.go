package main

import (
  "flag"
  "fmt"
)

type CmdFlags struct {
    Add    string
    Del    int
    Toggle int
    List   bool
}

func NewCmdFlags() *CmdFlags {
    cf := CmdFlags{}
    flag.StringVar(&cf.Add, "add", "", "Add a new todo")
    flag.IntVar(&cf.Del, "del", -1, "Delete todo by index")
    flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle todo by index")
    flag.BoolVar(&cf.List, "list", false, "List all todos")
    return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
    switch {
    case cf.List:
        todos.prints()
    case cf.Add != "":
        todos.add(cf.Add)
    case cf.Toggle != -1:
        todos.toggle(cf.Toggle)
    case cf.Del != -1:
        todos.deletes(cf.Del)
    default:
        fmt.Println("Invalid Command")
    }
}

package main
import (
  "fmt"

)

func main(){
  fmt.Println("\n")
  fmt.Println("==============================================================")
  fmt.Println("Welcome dude to Tamohar. This is not just a todo app.....")
  fmt.Println("It's your manageable journal and we will update this man.....")
  fmt.Println("Let's go  we will not give up....")
  fmt.Println("===============================================================")
  fmt.Println("\n")

  todos := Todos{}
  storage := NewStorage[Todos]("todo.json")
  storage.Load(&todos)

  cmdFlags := NewCmdFlags()
  flag.Parse()
  cmdFlags.Execute(&todos)
  storage.SaveData(todos)

 }

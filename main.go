package main
import (
  "fmt"
  "os"
)

func main(){
  fmt.Println("\n")
  fmt.Println("==============================================================")
  fmt.Println("Welcome dude to Tamohar. This is not just a todo app.....")
  fmt.Println("It's your manageable journal and we will update this man.....")
  fmt.Println("Let's go  we will not give up....")
  fmt.Println("===============================================================")
  fmt.Println("\n")

  fmt.Println("Enter day for the tasks. It could be Today, Tomorrow, Other")
  
  var days map[string]bool = map[string]bool{
    "Today": true, 
    "Tomorrow": true, 
    "Other": true,
  }

  var input string = os.Args[1]
  if !days[input] {
    fmt.Println("Be Specific man. Today, Tomorrow, Other(you will provide date here)")
    os.Exit(1)
  }
  // provide tthis handle switches boy
  fmt.Println(days[input])
  taskHandlers(input)

 }

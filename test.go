package main
//import "fmt"



func main(){
  todos := Todos{}
  todos.add("Learn Go fundamental")
  todos.add("Learn Mantra")
  todos.add("Learn Harmonica")
  todos.toggle(0)

   todos.prints()

 //fmt.Println(todos)
}

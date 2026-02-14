package main
import (   
  "fmt"
  "time"
  "bufio"
  "os"
  "math/rand"
  "strings"

)

type Tasks struct {
  id int
  Task_Info string
  isCompleted bool
  CreatedAt time.Time
  FinishedAt time.Time
}

func Input_Task () Tasks{
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Enter your Task: ")
  input,_ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  // random seed
  rand.Seed(time.Now().UnixNano())
  task_id :=  rand.Intn(1000) + 1

  return Tasks{
    id: task_id,
    Task_Info: input,
    isCompleted: false,
    CreatedAt: time.Now(),
    FinishedAt: time.Time{}, 
  }

}

func Assigned_Task(no_of_task int) {
  var all_task []Tasks

  for i:= 0; i<no_of_task ;i++ {
    current_tasks := Input_Task()
    all_task = append(all_task, current_tasks)

  }
  fmt.Println("all tasks are", all_task)
}

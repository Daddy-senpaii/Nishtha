package main
import (
  //"log"
  "fmt"
  "time"
  "bufio"
  "os"
  //"math/rand"
  "strings"
  //"encoding/json"

)

type Tasks struct {
  Id int  `json:"id"`
  Task_Info string `json:"task_info"`
  IsCompleted bool `json:"isCompleted"`
  CreatedAt time.Time `json:"createdAt"`
  FinishedAt time.Time `json:"finishedAt"`
}

func Input_Task () Tasks{
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Enter your Task: ")
  input,_ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  // random seed
  //rand.Seed(time.Now().UnixNano())
  //task_id :=  rand.Intn(1000) + 1
  //task_id := handle_id()

  tasks := Tasks{
    Id: NextID,
    Task_Info: input,
    IsCompleted: false,
    CreatedAt: time.Now(),
    FinishedAt: time.Time{}, 
  }
  NextID++
  return tasks
//  task_id += 1

}

func Assigned_Task(no_of_task int) {
  var all_task []Tasks
  handle_id()

  for i:= 0; i<no_of_task ;i++ {
    current_tasks := Input_Task()
    all_task = append(all_task, current_tasks)
  }

  receives_allTask(all_task)
}

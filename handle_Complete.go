package main 

import (
  "fmt"
  "log"
  "time"
)

func handle_Complete(task_id []int){
  fmt.Println("id's are ", task_id)
  all_task_data, err := file_data()
  if err != nil {
    log.Fatal(err)
  }
  
  for _, id := range task_id{
  task_index := get_task_location(all_task_data, id)

  if task_index == -1 {
    fmt.Println("Task not found: ", id)
    continue
  }

  //fmt.Printf("task_id: % T and data: %v", task_index, all_task_data[task_index])
  all_task_data[task_index].FinishedAt = time.Now()
  all_task_data[task_index].IsCompleted = true
  //fmt.Println("updated: ",all_task_data[task_index])
}

  all_taskData, err := struct_to_json(all_task_data)
  if err != nil {
    log.Fatal(err)
  }

  write_file(all_taskData)


}

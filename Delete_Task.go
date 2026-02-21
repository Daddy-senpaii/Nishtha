package main

import (
  "fmt"
  "log"
)

func get_task_location(current_data []Tasks, task_id int) int {
  
  var left, right int = 0, len(current_data)-1

  for left <= right {
    middle := (left + right)/2

    if current_data[middle].Id == task_id{
      return middle
    }
    if current_data[middle].Id < task_id {
      left = middle + 1
    }else {
      right = middle - 1
    }
    
  }
  return -1
}

func Delete_Task(id int){

  all_current_data, err := file_data() // getting data from file in json
  
  if err != nil {
    log.Fatal(err)
  }
  
 
  task_index := get_task_location(all_current_data, id) // finding index using binary search
  
  if task_index == -1 {

    fmt.Println("Task id not found boy")

  }else {
    all_current_data = append(all_current_data[:task_index], all_current_data[task_index+1:]...)

    encoded_data, err := struct_to_json(all_current_data)

    if err != nil {
      log.Fatal(err)
    }
        
    write_file(encoded_data) // write inside file man
    
  }
  
}

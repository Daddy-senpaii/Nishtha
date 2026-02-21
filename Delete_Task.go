package main

import (
  "fmt"
  "log"
)


func Delete_Task(id int){
  fmt.Println("Hey Delete here", id)
  all_current_data, err := file_data()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%T %d\n",all_current_data, len(all_current_data))
}

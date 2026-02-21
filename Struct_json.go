package main

import (
  "fmt"
  "encoding/json"
  "os"
  "log"
  "io"
)

func file_data()([]Tasks, error){

  fileData, err := os.OpenFile("all_task.json", os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  defer fileData.Close()

  data, err := io.ReadAll(fileData)

  if err != nil {
    return nil, err
  }

//  fmt.Printf("%v, %T", data, (data))
  all_current_data := json_to_struct(data)
  return all_current_data, nil
}

func receives_allTask(all_tasks []Tasks) {
    if _, err := os.Stat("all_task.json"); err == nil {
      fileData , err := os.OpenFile("all_task.json", os.O_RDWR, 0644)

      if err != nil {
        log.Fatal(err)
      }

      defer fileData.Close()

      data , err := io.ReadAll(fileData)
      if err != nil {
        log.Fatal(err)
     }

     all_current_tasks := json_to_struct(data)
   
    for _, value := range all_tasks {
      all_current_tasks = append(all_current_tasks, value)
    }

    fmt.Printf("%T\n",all_current_tasks)

    encoded_data, err := struct_to_json(all_current_tasks)
    if err != nil {
      log.Fatal(err)
    }

    err = os.WriteFile("all_task.json", encoded_data, 0644)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println("file is written successfully boy")

} else if os.IsNotExist(err) {

    data, err := struct_to_json(all_tasks)

    if err != nil {
      log.Fatal(err)
    }

    err = os.WriteFile("all_task.json", data, 0644)

    if err != nil {
      log.Fatal(err)
    }

    fmt.Println("boy file is made and json data is written successfully")


  }else{
    log.Fatal(err)
  }
}

func struct_to_json(all_task []Tasks) ([]byte, error){

  data, err := json.MarshalIndent(all_task, "", " ")
  return data, err
}

func json_to_struct(all_task []byte)([]Tasks){
  var tasks []Tasks
  err := json.Unmarshal(all_task, &tasks)
  if err != nil {
    log.Fatal(err)
  }
  return tasks
}

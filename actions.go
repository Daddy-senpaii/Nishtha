package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "errors"
  "strconv"
)

func command_handler(command string) (bool, error) {
  var command_look map[string]bool = map[string]bool{
    "add": true,
    "delete": true,
    "update":true,
    "get_all":true,
    "get_specific": true,
    "complete": true,
  }

  if _ , ok := command_look[command]; ok {
    return true, nil
  }
  return false, errors.New("Invalid Command boy.....: ");
}

func Actions(){
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Following things you can do in this....")

  fmt.Printf("1: Add Task using %s \n 2: Remove Task using %s \n 3: Update Task using %s \n 4: Get All Task using %s \n 5: Get Specific Task using %s \n 6: Complete Task using %s \n", "add", "delete", "update", "get_all", "get_specific", "complete" )

  fmt.Println("=====================================================================")

  fmt.Println("Enter command: ")
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)
  input_parts := strings.Split(input, " ")

  if len(input_parts) == 0{
    fmt.Println("No command entered")
    return 
  }
  
  command := input_parts[0]
  _ , err := command_handler(command)
  if err != nil {
    fmt.Println(err)
    return 
  }

  switch command {
    case "add" :
      fmt.Println("adding task")
      reader := bufio.NewReader(os.Stdin)
      fmt.Println("Enter the no of task you want to perform man. Be real man")
      text, _ := reader.ReadString('\n')
      text = strings.TrimSpace(text)

      no_of_task, err := strconv.Atoi(text)

      if err != nil {
        fmt.Println("Ivalid No")
        return 
  }
      Assigned_Task(no_of_task)
 
    case "update": 
      fmt.Println("updating task")
    case "delete":
      fmt.Println("removing task")
      fmt.Println("Enter id of task: ")
      reader := bufio.NewReader(os.Stdin)
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)

      task_id , err := strconv.Atoi(input)

      if err != nil {
        fmt.Println("Invalid id")
      }

      Delete_Task(task_id)

    case "get_all":
      all_tasks, err := file_data()
      if err != nil {
        fmt.Println("Having error man")
      }
      for _ , task := range all_tasks{
        fmt.Printf("\n")
        fmt.Println(task)
        fmt.Println("============================================")
      }
      
    case "get_specific":
      reader := bufio.NewReader(os.Stdin)
      fmt.Println("Please enter task id: you want to know")
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)

      task_id, err := strconv.Atoi(input)
      if err != nil {
        fmt.Println(err)
      }
      all_tasks, err := file_data()
      if err != nil {
        fmt.Println(err)
      }

      for _, task := range all_tasks{
        if task.Id == task_id {
          fmt.Println(task)
        }
      }
    case "complete":
      reader := bufio.NewReader(os.Stdin)
      fmt.Println("Please enter task id: ")
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      parts := strings.Fields(input)
      
      var taskIDs []int

      //task_id, err := strconv.Atoi(input)
      for _, value := range parts{
        id, err := strconv.Atoi(value)
        if err != nil {
          fmt.Println("Invalid number: ", value)
          break
        }
        taskIDs = append(taskIDs, id)
      }
      handle_Complete(taskIDs)
      }

    }
  

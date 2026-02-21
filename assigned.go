package main

//import ("fmt")

var NextID int = 0

func handle_id() int{
  all_data, err := file_data()

  if err != nil {
    NextID = 1
    return NextID
  }
  maxID := 0
  for _, value := range all_data {
    if value.Id > maxID {
      maxID = value.Id
    }
  }
  NextID = maxID + 1
  return NextID
}

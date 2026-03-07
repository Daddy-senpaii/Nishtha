package main
import (
  "errors"
  "time"
  "os"
  "strconv"
  "github.com/aquasecurity/table"
)

type Todo struct{
  Title string
  Completed bool
  CreatedAt time.Time
  CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) add(title string){
  todo := Todo{
    Title: title,
    Completed: false,
    CompletedAt: nil,
    CreatedAt: time.Now(),
  }
  *t = append(*t, todo)
}

func (todo *Todos) validateIndex(index int ) error {
  if index < 0 || index >= len(*todo){
    err := errors.New("invalid index")
    return err
  }
  return nil 
}

func(todo *Todos) deletes(index int) error {
  t := *todo
  if err := t.validateIndex(index); err != nil {
    return err
  }
  t = append(t[:index], t[index+1:]...)
  return nil 
}

func (todo *Todos) toggle(index int) error {
  t := *todo
  if err := t.validateIndex(index); err != nil {
    return err
  }

  if t[index].Completed {
    t[index].Completed = false
    t[index].CompletedAt = nil
  }else{
    now := time.Now()
    t[index].Completed = true
    t[index].CompletedAt = &now
  }
  return nil

}

func (todos *Todos) prints() {
  table := table.New(os.Stdout)
  table.SetRowLines(false)
  table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

  for index, t := range *todos {
    completed := "❌"
    completedAt := ""

    if t.Completed {
      completed = "✅"
      if t.CompletedAt != nil {
        completedAt = t.CompletedAt.Format(time.RFC1123)
      }
    }
    table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

  }
  table.Render()
}

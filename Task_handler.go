package main
import (
  "fmt"
  "os"
  "strconv"
  "time"
  "bufio"
  "strings"
  "errors"
)

func Date_Time_Handler(day string) (string, time.Weekday, error){
  now := time.Now()
  

  switch day {
    case "Today":
      date := fmt.Sprintf("%v-%v-%v", now.Day(), now.Month(), now.Year())
      weekday := now.Weekday()
      return date , weekday, nil

    case  "Tomorrow":
      tomorrow_date := now.AddDate(0, 0, 1)
      date := fmt.Sprintf("%v-%v-%v", tomorrow_date.Day(), tomorrow_date.Month(), tomorrow_date.Year())
      weekday := tomorrow_date.Weekday()
      return date, weekday, nil

    case "Other":
      reader := bufio.NewReader(os.Stdin)

      fmt.Print("Enter Day (1-31): ")
      dStr, _ := reader.ReadString('\n')
      dStr = strings.TrimSpace(dStr)
      dayInt, _ := strconv.Atoi(dStr)

      fmt.Print("Enter Month (1-12): ")
      mStr, _ := reader.ReadString('\n')
      mStr = strings.TrimSpace(mStr)
      monthInt, _ := strconv.Atoi(mStr)

      fmt.Print("Enter Year (e.g. 2026): ")
      yStr, _ := reader.ReadString('\n')
      yStr = strings.TrimSpace(yStr)
      yearInt,_ := strconv.Atoi(yStr)

      if yearInt < now.Year(){
        return "", now.Weekday(), errors.New("Please enter the valid year. You can't change past or live in past. So Be Smart. Enter Correct Year")
         
      }

      customDate := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.Local)
      return customDate.Format("02-01-2026"), customDate.Weekday(), nil

}
  return "", now.Weekday(), nil

  }

func taskHandlers(day string){
  date,weekday, err := Date_Time_Handler(day)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("Tasks for %s , %v \n", weekday, date) // update date here for right now

  Actions()  

}

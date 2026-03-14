package main

import (
  "flag"
  "fmt"
)
type ServerConfig struct {
  port int
  worker int
  env string
}

type DockerConfig struct {
  image string
  porty int
}

func main() {
  var server ServerConfig
  var docker DockerConfig

  flag.StringVar(&docker.image, "image", "panty", "image name string")
  flag.IntVar(&docker.porty, "porty", 2020, "integer of port name")

  flag.IntVar(&server.port, "port", 8000, "Server Port")
  flag.IntVar(&server.worker,"worker",  2, "Workers need")
  flag.StringVar(&server.env, "env", "",  "enviornment variable")

  var age = flag.Int("age",20, "An age integer")
  var name string
  flag.StringVar(&name,"name", "", "A name of string")

  var a int
  flag.IntVar(&a,"ages", 10, "An int integer")
  flag.Parse()

  fmt.Printf("package is %s and port is %d\n", docker.image, docker.porty)
  fmt.Println(server)
  fmt.Println(*age)
  fmt.Println(name)
  fmt.Println(a)
}

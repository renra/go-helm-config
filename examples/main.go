package main

import (
  "os"
  "fmt"
  "app/config"
)

func main() {
  os.Setenv("HERO_NAME", "Jon")
  conf := config.Load("charts/go-helm-logger/env", "testing")

  fmt.Println(conf.Get("non_existing_key"))
  fmt.Println(conf.GetString("non_existing_key"))

  fmt.Println(conf.Get("height"))
  fmt.Println(conf.GetString("height"))
  fmt.Println(conf.GetString("weight"))

  stagingConf := config.Load("charts/go-helm-logger/env", "staging")
  fmt.Println(stagingConf.GetString("height"))

  fmt.Println(stagingConf.GetString("gopath"))
  fmt.Println(stagingConf.GetString("hero_name"))
}


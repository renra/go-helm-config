package main

import (
  "os"
  "fmt"
  "app/config"
)

func main() {
  os.Setenv("HERO_NAME", "Jon")

  conf, err := config.Load("charts/go-helm-logger/env", "testing")
  fmt.Println(err)

  non_existing_key, err := conf.Get("non_existing_key")
  fmt.Println(non_existing_key)
  fmt.Println(err)

  non_existing_key, err = conf.GetString("non_existing_key")
  fmt.Println(non_existing_key)
  fmt.Println(err)

  height, err := conf.Get("height")
  fmt.Println(height)
  fmt.Println(err)

  height, err = conf.GetString("height")
  fmt.Println(height)
  fmt.Println(err)

  weight, err := conf.GetString("weight")
  fmt.Println(weight)
  fmt.Println(err)

  stagingConf, err := config.Load("charts/go-helm-logger/env", "staging")
  fmt.Println(err)

  height, err = stagingConf.GetString("height")
  fmt.Println(height)
  fmt.Println(err)

  gopath, err := stagingConf.GetString("gopath")
  fmt.Println(gopath)
  fmt.Println(err)

  heroName, err := stagingConf.GetString("hero_name")
  fmt.Println(heroName)
  fmt.Println(err)
}


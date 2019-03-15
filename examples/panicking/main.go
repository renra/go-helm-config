package main

import (
  "os"
  "fmt"
  "app/config"
)

func main() {
  os.Setenv("HERO_NAME", "Jon")

  conf := config.LoadP("charts/go-helm-logger/env", "testing")

  height := conf.GetP("height")
  fmt.Println(height)

  height = conf.GetStringP("height")
  fmt.Println(height)

  weight := conf.GetStringP("weight")
  fmt.Println(weight)

  stagingConf := config.LoadP("charts/go-helm-logger/env", "staging")

  height = stagingConf.GetStringP("height")
  fmt.Println(height)

  gopath := stagingConf.GetStringP("gopath")
  fmt.Println(gopath)

  heroName := stagingConf.GetStringP("hero_name")
  fmt.Println(heroName)

  //non_existing_key, err := conf.Get("non_existing_key")
  //fmt.Println(non_existing_key)
  //fmt.Println(err)
}



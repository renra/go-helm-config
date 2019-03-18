package main

import (
  "os"
  "fmt"
  "app/config"
)

func main() {
  os.Setenv("HERO_NAME", "Jon")

  conf := config.LoadP("charts/go-helm-config/env", "testing")

  height := conf.GetP("height")
  fmt.Println(height)

  height = conf.GetStringP("height")
  fmt.Println(height)

  weight := conf.GetStringP("weight")
  fmt.Println(weight)

  weightInt := conf.GetIntP("weight")
  fmt.Println(weightInt)

  weightFloat := conf.GetFloatP("weight")
  fmt.Println(weightFloat)

  flag := conf.GetBoolP("flag")
  fmt.Println(flag)

  stagingConf := config.LoadP("charts/go-helm-config/env", "staging")

  height = stagingConf.GetStringP("height")
  fmt.Println(height)

  gopath := stagingConf.GetStringP("gopath")
  fmt.Println(gopath)

  heroName := stagingConf.GetStringP("hero_name")
  fmt.Println(heroName)
}



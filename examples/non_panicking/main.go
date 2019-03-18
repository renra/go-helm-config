package main

import (
  "os"
  "fmt"
  "app/config"
)

func main() {
  os.Setenv("HERO_NAME", "Jon")

  conf, fatalErr, ignorableErr := config.Load("charts/go-helm-config/env", "testing")
  fmt.Println(fmt.Sprintf("Fatal: %v", fatalErr))
  fmt.Println(fmt.Sprintf("Ignorable: %v", ignorableErr))

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

  weightBool, err := conf.GetBool("weight")
  fmt.Println(weightBool)
  fmt.Println(err)

  weightFloat, err := conf.GetFloat("weight")
  fmt.Println(weightFloat)
  fmt.Println(err)

  weightInt, err := conf.GetInt("weight")
  fmt.Println(weightInt)
  fmt.Println(err)

  stagingConf, fatalErr, ignorableErr := config.Load("charts/go-helm-config/env", "staging")
  fmt.Println(fmt.Sprintf("Fatal: %v", fatalErr))
  fmt.Println(fmt.Sprintf("Ignorable: %v", ignorableErr))

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


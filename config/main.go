package config

import (
  "os"
  "fmt"
  "strings"
  "gopkg.in/yaml.v2"
  "github.com/gobuffalo/packr/v2"
)

type ConfigData map[string]interface{}

type Config struct {
  Data ConfigData
}

func (c *ConfigData) Branch(key string) *ConfigData {
  result := make(ConfigData)

  branch, ok := (*c)[key]

  if ok == false {
    panic(fmt.Sprintf("Could not read key: %s", key))
  }

  if branch == nil {
    branch = make(map[interface{}]interface{})
  }

  for k, v := range branch.(map[interface{}]interface{}) {
    result[k.(string)] = v
  }

  return &result
}

func (c *Config) Get(key string) interface{} {
  return c.Data[key]
}

func (c *Config) GetString(key string) string {
  value := c.Get(key)

  if value == nil {
    return ""
  }

  strValue, ok := value.(string)

  if ok != true {
    panic(fmt.Sprintf("Config value %s not strigifyable", strValue))
  }

  return strValue;
}

// Path is split into two to prevent creating boxes with unnecessary files
//  for example packs.New("Whatever", "./") would compile all files in the project
//  and include it in the binary
func loadConfig(pathToDir string, fileName string) (*ConfigData, error) {
  box := packr.New(fmt.Sprintln("Config - %s", pathToDir), pathToDir)

  configInYaml, err := box.FindString(fileName)

  if err != nil {
    return nil, err
  }

  configData := ConfigData{}
  err = yaml.Unmarshal([]byte(configInYaml), &configData)

  if err != nil {
    panic(err)
  }

  return configData.Branch("env_vars"), nil
}

func Load(basePath string, env string) *Config {
  configData, err := loadConfig(basePath, "values.yaml")

  if err != nil {
    panic(err)
  }

  // This is practically not necessary because the app will have all of these in env vars
  //  whenever it is running in helm. However it wouldn't be the case if running on localhost
  //  on staging, testing or production mode (for whatever reasons)
  path := fmt.Sprintf("%s/%s/", basePath, env)
  envConfig, _ := loadConfig(path, "values.yaml")

  if envConfig != nil {
    for k, v := range *envConfig {
      (*configData)[k] = v
    }
  }

  for _, envVar := range os.Environ() {
    pair := strings.Split(envVar, "=")

    k := strings.ToLower(pair[0])
    v := pair[1]

    (*configData)[k] = v
  }

  return &Config{Data: (*configData)}
}


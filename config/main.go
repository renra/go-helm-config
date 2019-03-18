package config

import (
  "fmt"
  "github.com/renra/go-errtrace/errtrace"
  "github.com/renra/go-yaml-config/config"
)

type Config struct {
  yamlConfig *config.Config
}

func (c *Config) C() *config.Config {
  return c.yamlConfig
}

func (c *Config) Get(key string) (interface{}, *errtrace.Error) {
  return c.yamlConfig.Get(key)
}

func (c *Config) GetP(key string) interface{} {
  return c.yamlConfig.GetP(key)
}

func (c *Config) GetString(key string) (string, *errtrace.Error) {
  return c.yamlConfig.GetString(key)
}

func (c *Config) GetStringP(key string) string {
  return c.yamlConfig.GetStringP(key)
}

func (c *Config) GetInt(key string) (int, *errtrace.Error) {
  return c.yamlConfig.GetInt(key)
}

func (c *Config) GetIntP(key string) int {
  return c.yamlConfig.GetIntP(key)
}

func (c *Config) GetFloat(key string) (float64, *errtrace.Error) {
  return c.yamlConfig.GetFloat(key)
}

func (c *Config) GetFloatP(key string) float64 {
  return c.yamlConfig.GetFloatP(key)
}

func (c *Config) GetBool(key string) (bool, *errtrace.Error) {
  return c.yamlConfig.GetBool(key)
}

func (c *Config) GetBoolP(key string) bool {
  return c.yamlConfig.GetBoolP(key)
}

func Load(basePath string, env string) (*Config, *errtrace.Error, *errtrace.Error) {
  c1, err1 := config.LoadSection(fmt.Sprintf("%s/values.yaml", basePath), "env_vars")
  c2, err2 := config.LoadSection(fmt.Sprintf("%s/%s/values.yaml", basePath, env), "env_vars")

  if err1 != nil {
    if err2 != nil {
      return nil, errtrace.Wrap(err1), errtrace.Wrap(err2)
    } else {
      return nil, errtrace.Wrap(err1), nil
    }
  }

  if err2 != nil {
    return &Config{yamlConfig: c1.MergeWithEnvVars()}, nil, errtrace.Wrap(err2)
  } else {
    return &Config{yamlConfig: c1.Merge(c2).MergeWithEnvVars()}, nil, nil
  }
}

func LoadP(basePath string, env string) *Config {
  conf, fatalErr, _ := Load(basePath, env)

  if fatalErr != nil {
    panic(fatalErr)
  }

  return conf
}

package config

import (
  "fmt"
  "github.com/renra/go-yaml-config/config"
)

type Config struct {
  yamlConfig *config.Config
}

func (c *Config) C() *config.Config {
  return c.yamlConfig
}

func (c *Config) Get(key string) (interface{}, error) {
  return c.yamlConfig.Get(key)
}

func (c *Config) GetP(key string) interface{} {
  return c.yamlConfig.GetP(key)
}

func (c *Config) GetString(key string) (string, error) {
  return c.yamlConfig.GetString(key)
}

func (c *Config) GetStringP(key string) string {
  return c.yamlConfig.GetStringP(key)
}

func Load(basePath string, env string) (*Config, error) {
  c1, err := config.LoadSection(fmt.Sprintf("%s/values.yaml", basePath), "env_vars")

  if err != nil {
    panic(err)
  }

  c2, err := config.LoadSection(fmt.Sprintf("%s/%s/values.yaml", basePath, env), "env_vars")

  if err != nil {
    panic(err)
  }

  return &Config{yamlConfig: c1.Merge(c2).MergeWithEnvVars()}, nil
}

func LoadP(basePath string, env string) *Config {
  conf, err := Load(basePath, env)

  if err != nil {
    panic(err)
  }

  return conf
}

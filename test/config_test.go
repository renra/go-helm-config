package main

import (
  "os"
  "fmt"
  "testing"
  "io/ioutil"
  "app/config"
  "gopkg.in/yaml.v2"
  "github.com/stretchr/testify/suite"
)

var mainFileName string = "./test/values.yaml"
var primaryWidth int = 200
var primaryHeight float64 = 200.5
var primaryLength int = 200
var numbers [3]string = [3]string{"one", "two", "three"}
var primaryIsAwesome bool = false
var primaryIsTerrible bool = true

var env = "whatever"
var secondaryFileName string = fmt.Sprintf("./test/%s/values.yaml", env)
var secondaryWidth int = 400
var secondaryHeight int = 400

var tertiaryHeight int = 600
var heroName string = "Jon"
var section string = "env_vars"
var tertiaryIsAwesome bool = true
var tertiaryIsTerrible bool = false

func writeYaml(path string, data map[string]interface{}) {
  contents, err := yaml.Marshal(&data)

  if err != nil {
    panic(err)
  }

  err = ioutil.WriteFile(path, contents, 0644)

  if err != nil {
    panic(err)
  }
}

func writePrimaryConfigFile() {
  writeYaml(mainFileName, map[string]interface{}{
    section: map[string]interface{}{
      "width": primaryWidth,
      "height": primaryHeight,
      "length": primaryLength,
      "numbers": numbers,
      "is_awesome": primaryIsAwesome,
      "is_terrible": primaryIsTerrible,
    },
  })
}

func writeSecondaryConfigFile() {
  err := os.Mkdir(fmt.Sprintf("./test/%s", env), 0644)

  if err != nil {
    panic(err)
  }

  writeYaml(secondaryFileName, map[string]interface{}{
    section: map[string]interface{}{
      "width": secondaryWidth,
      "height": secondaryHeight,
    },
  })
}

func setupEnvVars() {
  os.Setenv("HERO_NAME", heroName)
  os.Setenv("HEIGHT", fmt.Sprintf("%d", tertiaryHeight))

  os.Setenv("IS_AWESOME", fmt.Sprintf("%t", tertiaryIsAwesome))
  os.Setenv("IS_TERRIBLE", fmt.Sprintf("%t", tertiaryIsTerrible))
}

type ConfigSuite struct {
  suite.Suite
}

func TestConfigSuite(t *testing.T) {
  suite.Run(t, new(ConfigSuite))
}

func (c *ConfigSuite) TearDownTest() {
  os.Remove(mainFileName)
  os.Remove(secondaryFileName)
  os.Remove(fmt.Sprintf("./test/%s", env))
}

func (c *ConfigSuite) TestLoad() {
  writePrimaryConfigFile()
  writeSecondaryConfigFile()
  setupEnvVars()

  config, fatalError, ignorableError := config.Load("test", env)
  if fatalError != nil {
    panic("Could not load main config file")
  }

  if ignorableError != nil {
    panic("Could not load overrides")
  }

  expectedWidth := fmt.Sprintf("%d", secondaryWidth)
  widthFromConfig, err := config.GetString("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: width")
  }

  expectedHeight := fmt.Sprintf("%d", tertiaryHeight)
  heightFromConfig, err := config.GetString("height")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %s, got %s", expectedHeight, heightFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: height")
  }

  expectedLength := fmt.Sprintf("%d", primaryLength)
  lengthFromConfig, err := config.GetString("length")

  if lengthFromConfig != expectedLength {
    c.T().Errorf("Expected %s, got %s", expectedLength, lengthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: length")
  }

  expectedIsAwesome := fmt.Sprintf("%t", tertiaryIsAwesome)
  isAwesomeFromConfig, err := config.GetString("is_awesome")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %s, got %s", expectedIsAwesome, isAwesomeFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_awesome")
  }

  expectedIsTerrible := fmt.Sprintf("%t", tertiaryIsTerrible)
  isTerribleFromConfig, err := config.GetString("is_terrible")

  if isTerribleFromConfig != expectedIsTerrible {
    c.T().Errorf("Expected %s, got %s", expectedIsTerrible, isTerribleFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_terrible")
  }

  expectedValue := ""
  unexistingValue, err := config.GetString("unexisting")

  if unexistingValue != expectedValue {
    c.T().Errorf("Expected %s, got %s", expectedValue, unexistingValue)
  }

  if err == nil {
    c.T().Errorf("Expected not to find key: unexisting")
  }
}

func (c *ConfigSuite) TestLoadUnexistingSecondaryFile() {
  writePrimaryConfigFile()
  setupEnvVars()

  config, fatalError, ignorableError := config.Load("test", env)
  if fatalError != nil {
    panic("Could not load main config file")
  }

  if ignorableError == nil {
    c.T().Errorf("Expected to see error here")
  }

  expectedWidth := fmt.Sprintf("%d", primaryWidth)
  widthFromConfig, err := config.GetString("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: width")
  }

  expectedHeight := fmt.Sprintf("%d", tertiaryHeight)
  heightFromConfig, err := config.GetString("height")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %s, got %s", expectedHeight, heightFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: height")
  }

  expectedLength := fmt.Sprintf("%d", primaryLength)
  lengthFromConfig, err := config.GetString("length")

  if lengthFromConfig != expectedLength {
    c.T().Errorf("Expected %s, got %s", expectedLength, lengthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: length")
  }

  expectedIsAwesome := fmt.Sprintf("%t", tertiaryIsAwesome)
  isAwesomeFromConfig, err := config.GetString("is_awesome")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %s, got %s", expectedIsAwesome, isAwesomeFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_awesome")
  }

  expectedIsTerrible := fmt.Sprintf("%t", tertiaryIsTerrible)
  isTerribleFromConfig, err := config.GetString("is_terrible")

  if isTerribleFromConfig != expectedIsTerrible {
    c.T().Errorf("Expected %s, got %s", expectedIsTerrible, isTerribleFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_terrible")
  }

  expectedValue := ""
  unexistingValue, err := config.GetString("unexisting")

  if unexistingValue != expectedValue {
    c.T().Errorf("Expected %s, got %s", expectedValue, unexistingValue)
  }

  if err == nil {
    c.T().Errorf("Expected not to find key: unexisting")
  }
}

func (c *ConfigSuite) TestLoadUnexistingPrimaryFile() {
  _, fatalError, ignorableError := config.Load("test", env)
  if fatalError == nil {
    c.T().Errorf("Expected to see error here")
  }

  if ignorableError == nil {
    c.T().Errorf("Expected to see error here")
  }
}

func (c *ConfigSuite) TestLoadP() {
  writePrimaryConfigFile()
  writeSecondaryConfigFile()
  setupEnvVars()

  config := config.LoadP("test", env)

  expectedWidth := fmt.Sprintf("%d", secondaryWidth)
  widthFromConfig, err := config.GetString("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: width")
  }

  expectedHeight := fmt.Sprintf("%d", tertiaryHeight)
  heightFromConfig, err := config.GetString("height")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %s, got %s", expectedHeight, heightFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: height")
  }

  expectedLength := fmt.Sprintf("%d", primaryLength)
  lengthFromConfig, err := config.GetString("length")

  if lengthFromConfig != expectedLength {
    c.T().Errorf("Expected %s, got %s", expectedLength, lengthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: length")
  }

  expectedIsAwesome := fmt.Sprintf("%t", tertiaryIsAwesome)
  isAwesomeFromConfig, err := config.GetString("is_awesome")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %s, got %s", expectedIsAwesome, isAwesomeFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_awesome")
  }

  expectedIsTerrible := fmt.Sprintf("%t", tertiaryIsTerrible)
  isTerribleFromConfig, err := config.GetString("is_terrible")

  if isTerribleFromConfig != expectedIsTerrible {
    c.T().Errorf("Expected %s, got %s", expectedIsTerrible, isTerribleFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_terrible")
  }

  expectedValue := ""
  unexistingValue, err := config.GetString("unexisting")

  if unexistingValue != expectedValue {
    c.T().Errorf("Expected %s, got %s", expectedValue, unexistingValue)
  }

  if err == nil {
    c.T().Errorf("Expected not to find key: unexisting")
  }
}

func (c *ConfigSuite) TestLoadPUnexistingSecondaryFile() {
  writePrimaryConfigFile()
  setupEnvVars()

  config := config.LoadP("test", env)

  expectedWidth := fmt.Sprintf("%d", primaryWidth)
  widthFromConfig, err := config.GetString("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: width")
  }

  expectedHeight := fmt.Sprintf("%d", tertiaryHeight)
  heightFromConfig, err := config.GetString("height")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %s, got %s", expectedHeight, heightFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: height")
  }

  expectedLength := fmt.Sprintf("%d", primaryLength)
  lengthFromConfig, err := config.GetString("length")

  if lengthFromConfig != expectedLength {
    c.T().Errorf("Expected %s, got %s", expectedLength, lengthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: length")
  }

  expectedIsAwesome := fmt.Sprintf("%t", tertiaryIsAwesome)
  isAwesomeFromConfig, err := config.GetString("is_awesome")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %s, got %s", expectedIsAwesome, isAwesomeFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_awesome")
  }

  expectedIsTerrible := fmt.Sprintf("%t", tertiaryIsTerrible)
  isTerribleFromConfig, err := config.GetString("is_terrible")

  if isTerribleFromConfig != expectedIsTerrible {
    c.T().Errorf("Expected %s, got %s", expectedIsTerrible, isTerribleFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected to find key: is_terrible")
  }

  expectedValue := ""
  unexistingValue, err := config.GetString("unexisting")

  if unexistingValue != expectedValue {
    c.T().Errorf("Expected %s, got %s", expectedValue, unexistingValue)
  }

  if err == nil {
    c.T().Errorf("Expected not to find key: unexisting")
  }
}

func (c *ConfigSuite) TestLoadPUnexistingPrimaryFile() {
  defer func() {
    r := recover()

    if r == nil {
      c.T().Errorf("Expected to see error here")
    }
  }()

  config.LoadP("test", env)
}

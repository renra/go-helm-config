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
      "LENGTH": primaryLength,
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
      "WIDTH": secondaryWidth,
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

func (c *ConfigSuite) TestSet() {
  writePrimaryConfigFile()
  writeSecondaryConfigFile()
  setupEnvVars()

  config := config.LoadP("test", env)

  expectedWidth := 1600
  config.Set("width", expectedWidth)

  width, err := config.Get("width")

  if err != nil {
    c.T().Errorf("Expected to find key: width")
  }

  if width.(int) != expectedWidth {
    c.T().Errorf("Expected %v, got %v", expectedWidth, width)
  }
}

func (c *ConfigSuite) TestGetString() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedWidth := fmt.Sprintf("%d", primaryWidth)
  widthFromConfig, err := config.GetString("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected not to see error here")
  }
}

func (c *ConfigSuite) TestGetStringUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedWidth := ""
  widthFromConfig, err := config.GetString("unexisting")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }

  if err == nil {
    c.T().Errorf("Expected to see error here")
  }
}

func (c *ConfigSuite) TestGetStringP() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedWidth := fmt.Sprintf("%d", primaryWidth)
  widthFromConfig := config.GetStringP("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %s, got %s", expectedWidth, widthFromConfig)
  }
}

func (c *ConfigSuite) TestGetStringPUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  defer func() {
    r := recover()

    if r == nil {
      c.T().Errorf("Expected to see error here")
    }
  }()

  config.GetStringP("unexisting")
}

func (c *ConfigSuite) TestGetInt() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedWidth := primaryWidth
  widthFromConfig, err := config.GetInt("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %d, got %d", expectedWidth, widthFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected not to see error here")
  }
}

func (c *ConfigSuite) TestGetIntUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedWidth := 0
  widthFromConfig, err := config.GetInt("unexisting")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %d, got %d", expectedWidth, widthFromConfig)
  }

  if err == nil {
    c.T().Errorf("Expected to see error here")
  }
}

func (c *ConfigSuite) TestGetIntP() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedWidth := primaryWidth
  widthFromConfig := config.GetIntP("width")

  if widthFromConfig != expectedWidth {
    c.T().Errorf("Expected %d, got %d", expectedWidth, widthFromConfig)
  }
}

func (c *ConfigSuite) TestGetIntPUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  defer func() {
    r := recover()

    if r == nil {
      c.T().Errorf("Expected to see error here")
    }
  }()

  config.GetIntP("unexisting")
}

func (c *ConfigSuite) TestGetFloat() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedHeight := primaryHeight
  heightFromConfig, err := config.GetFloat("height")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %f, got %f", expectedHeight, heightFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected not to see error here")
  }
}

func (c *ConfigSuite) TestGetFloatUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedHeight := 0.0
  heightFromConfig, err := config.GetFloat("unexisting")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %f, got %f", expectedHeight, heightFromConfig)
  }

  if err == nil {
    c.T().Errorf("Expected to see error here")
  }
}

func (c *ConfigSuite) TestGetFloatP() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedHeight := primaryHeight
  heightFromConfig := config.GetFloatP("height")

  if heightFromConfig != expectedHeight {
    c.T().Errorf("Expected %f, got %f", expectedHeight, heightFromConfig)
  }
}

func (c *ConfigSuite) TestGetFloatPUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  defer func() {
    r := recover()

    if r == nil {
      c.T().Errorf("Expected to see error here")
    }
  }()

  config.GetFloatP("unexisting")
}

func (c *ConfigSuite) TestGetBool() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedIsAwesome := primaryIsAwesome
  isAwesomeFromConfig, err := config.GetBool("is_awesome")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %t, got %t", expectedIsAwesome, isAwesomeFromConfig)
  }

  if err != nil {
    c.T().Errorf("Expected not to see error here")
  }
}

func (c *ConfigSuite) TestGetBoolUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedIsAwesome := false
  isAwesomeFromConfig, err := config.GetBool("unexisting")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %t, got %t", expectedIsAwesome, isAwesomeFromConfig)
  }

  if err == nil {
    c.T().Errorf("Expected to see error here")
  }
}

func (c *ConfigSuite) TestGetBoolP() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  expectedIsAwesome := primaryIsAwesome
  isAwesomeFromConfig := config.GetBoolP("is_awesome")

  if isAwesomeFromConfig != expectedIsAwesome {
    c.T().Errorf("Expected %t, got %t", expectedIsAwesome, isAwesomeFromConfig)
  }
}

func (c *ConfigSuite) TestGetBoolPUnexistingKey() {
  writePrimaryConfigFile()

  config := config.LoadP("test", env)

  defer func() {
    r := recover()

    if r == nil {
      c.T().Errorf("Expected to see error here")
    }
  }()

  config.GetBoolP("unexisting")
}

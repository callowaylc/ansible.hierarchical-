package hieransible

import (
  "io/ioutil"
  "os"
  yaml "gopkg.in/yaml.v2"
)

type Config struct {
  Stack []string
}
var c Config = Config{}

func InitConfig() {
  contents, err := ioutil.ReadFile("./config.yml")
  if err != nil {
    Logs("Failed to load config", LogFields{
      "error": err,
    })
    os.Exit(1)
  }
  yaml.Unmarshal([]byte(contents), c)
}

func GetConfig() Config {
  return c
}

package configuration

import (
  "encoding/json"
  "os"
  "io/ioutil"

  mlg "github.com/mattrax/mattrax/internal/logging" //Mattrax Logging
)

var (
  log = mlg.GetLogger()
  config = Config{}
)

func init() {
  if configFile, err := os.Open("config.json"); os.IsNotExist(err) {
    json, err := json.MarshalIndent(newConfig(), "", "  ")
    if err != nil { log.Fatal("Error Generating The Config File:", err) }
    if err := ioutil.WriteFile("config.json", json, 0644); err != nil { log.Fatal("Error Saving The New Config File To './config.json'") }
    log.Warning("A New Config Was Created. Please Populate The Correct Information Before Starting Mattrax Again.")
    return
  } else if err != nil {
    log.Fatal("Error Loading The Config File:", err)
  } else {
    if err := json.NewDecoder(configFile).Decode(&config); err != nil { log.Fatal("Error Parsing The Config File:", err) }
  }
}

func GetConfig() Config { return config }

func newConfig() Config {
  return Config{
    Name: "Acme Inc",
    Domain: "mdm.acme.com",
    Verbose: false,
    LogFile: "data/log.txt",
    Port: 8000,
    Database: "postgres://postgres:@postgres/postgres",
  }
}

type Config struct {
  Name string `json:"name"`
  Domain string `json:"domain"`
  Verbose bool `json:"verbose"`
  LogFile string `json:"logFile"`
  Port int `json:"port"`
  Database string `json:"database"`
}

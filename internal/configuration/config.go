package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log" //Uses This Instead of My Logger Because of Import Cycle Error
	"os"
)

var config = Config{} // The Configuration

// TODO: Go Doc
func init() {
	if configFile, err := os.Open("config.json"); os.IsNotExist(err) {
		json, err := json.MarshalIndent(newConfig(), "", "  ")
		if err != nil {
			log.Fatal("Error Generating The Config File:", err)
		}
		if err := ioutil.WriteFile("config.json", json, 0644); err != nil {
			log.Fatal("Error Saving The New Config File To './config.json'")
		}
		log.Println("A New Config Was Created. Please Populate The Correct Information Before Starting Mattrax Again.")
		os.Exit(0)
	} else if err != nil {
		log.Fatal("Error Loading The Config File:", err)
	} else {
		if err := json.NewDecoder(configFile).Decode(&config); err != nil {
			log.Fatal("Error Parsing The Config File:", err)
		}
	}
}

// TODO: Go Doc
func GetConfig() Config { return config }

// TODO: Go Doc
func newConfig() Config {
	return Config{
		Name:     "Acme Inc",
		Domain:   "mdm.acme.com",
		EEDomain: "enterpriseenrollment.acme.com",
		Verbose:  false,
		LogFile:  "data/log.txt",
		APNScert: "data/PushCert.p123",
		APNSProduction: true,
		APNSpass: "password",
		Port:     8000,
		Database: "postgres://postgres:@postgres/postgres",
	}
}

// TODO: Go Doc
type Config struct {
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	EEDomain string `json:"EEDomain"`
	Verbose  bool   `json:"verbose"`
	LogFile  string `json:"logFile"`
	APNScert string `json:"APNScert"`
	APNSProduction bool `json:"APNSProduction"`
	APNSpass string `json:"APNSpass"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

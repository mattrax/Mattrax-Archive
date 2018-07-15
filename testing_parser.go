package main

import (
  "log"
  "github.com/groob/plist" //Plist Parsing
)

type Command struct {
  RequestType string `plist:"RequestType,notnull" json:"request_type"`
  InstallApplication
}

type InstallApplication struct {
	ITunesStoreID   int `plist:"iTunesStoreID,omitempty"`
}

func main() {
  payload := Command{
    RequestType: "InstallApplication",
    InstallApplication: InstallApplication{
      ITunesStoreID: 640199958,
    },
  }

  plistCmd, err := plist.MarshalIndent(payload, "\t")
	if err != nil { log.Println(err); return }
  log.Println(string(plistCmd))
}

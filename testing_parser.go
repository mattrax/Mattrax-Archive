package main

import (
  "log"
  "github.com/groob/plist" //Plist Parsing

  //"reflect"
  "encoding/json"
  "gopkg.in/oleiade/reflections.v1"
)

type Command struct {
  RequestType string `plist:"RequestType,notnull" json:"request_type"`
  InstallApplication
}

type Command2 struct {
  InstallApplication
}


type InstallApplication struct {
	ITunesStoreID   int `plist:"iTunesStoreID,omitempty"`
}

type Database struct {
  InstallApplication
	//ITunesStoreID int // TODO: Is This int64
  //ManagementFlags string
}








type Command3 struct {
  RequestType string `plist:"RequestType"`
  TestingJSON
}

type TestingJSON struct {
  InstallApplication
    //ITunesStoreID int `json:"ITunesStoreID"` //TODO .omit empty
    //ManagementFlags  string `json:"ManagementFlags"` //TODO .omit empty
}

func main() {
  policyOptions := `{

  "ITunesStoreID" : 640199958
}` //"ManagementFlags" : 4,

  /*current := Command3{
    RequestType: "InstallApplication",
  }*/

  var options Command3
	if err := json.Unmarshal([]byte(policyOptions), &options); err != nil {
		log.Println(err); return
	}
  options.RequestType = "InstallApplication"   //"RequestType": "InstallApplication",
  //RequestType: "InstallApplication",


  log.Println(options)

  plistCmd, err := plist.MarshalIndent(options, "\t")
	if err != nil { log.Println(err); return }
  log.Println(string(plistCmd))


}
















func old() {
  /*payload := Command{
    RequestType: "InstallApplication",
    InstallApplication: InstallApplication{
      ITunesStoreID: 640199958,
    },
  }*/

  dbPolicy := Database{
    InstallApplication: InstallApplication{
      ITunesStoreID: 640199958,
    },
  }

  RequestType := "InstallApplication"
  //iTunesStoreID := 640199958



  payload := Command{
    RequestType: RequestType,
  }

  _ = reflections.SetField(&payload, RequestType, dbPolicy)/*InstallApplication{
    ITunesStoreID: iTunesStoreID,
  })*/






  plistCmd, err := plist.MarshalIndent(payload, "\t")
	if err != nil { log.Println(err); return }
  log.Println(string(plistCmd))
}

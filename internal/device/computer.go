package main //package device

import (
  "fmt"
  "github.com/go-pg/pg"
)

/*test := Computer{
  Name: "Testing Qwerty",
}
fmt.Println(test.Testing())
fmt.Println(test.GetName())

fmt.Println("")

aPC := WindowsComputer{
  Computer: Computer{
    Name: "Testing Qwerty",
  },
  Testing123: "This Is A Windows Computer",
}
fmt.Println(aPC.Computer.GetName())
fmt.Println(aPC.Testing123)*/

type WindowsComputer struct {
    Computer
    Testing123 string
}


type Device struct { /* TEMP */
	TableName      struct{}       `sql:"devices"`
	UUID           string         `sql:"uuid,pk"`

	/*DeviceState    int            `sql:"DeviceState,notnull"`
	DeviceDetails  DeviceDetails  `sql:"DeviceDetails,notnull"`
	DeviceTokens   DeviceTokens   `sql:"DeviceTokens,notnull"`
	DevicePolicies DevicePolicies `sql:"DevicePolicies,notnull"`*/
}

//Use PG .Table("author_books")
/*var devices []Device
if err := pgdb.Model(&devices).Select(); err != nil {
  fmt.Println(err)
  return
}

for _, device := range devices {
  fmt.Println(device.UDID)
}*/

func main() {
  pgdb := pg.Connect(&pg.Options{
      User: "oscar.beaumont",
      Database: "mattrax2",
  })
  defer pgdb.Close()

  var device Device
	if err := pgdb.Model(&device).Where("uuid = ?", "877914b242e969ee82abf93537b16db3a2441ae9").Select(); err != nil {
		fmt.Println(err)
    return
	}
  fmt.Println(device.UUID)






  /*comp := Computer{}

  fmt.Println(comp)
  fmt.Println(comp.Name)*/
}












type Computer struct {
  Name string //`plist:"Name,omitempty"`
  UUID string //`plist:"uuid,omitempty"`
  DeviceState int //`plist:"DeviceState,omitempty"`

  //Device Information
  //Network Interfaces
  //Operating System Informaion -> Versions
  //Installed Applications -> Maybe Device Specific
}

func (comp Computer) FillDeviceInformation() string {
  return "Testing123"
}

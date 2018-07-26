package main //package device

import (
  "fmt"
  "github.com/go-pg/pg"
)

/* Database Interface */
type DeviceInformation struct {
  DisplayName string `sql:"DisplayName"`
  // Enrollment State (Genericly Designed)
  // OS Information
  // /
}

type Device struct {
	TableName struct{} `sql:"devices"`
	UUID string `sql:"uuid,pk"`
  DeviceInformation DeviceInformation `sql:"DeviceInformation"`
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




type WindowsComputer struct {
    Computer
    Testing123 string
}

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

  // Create The Computer
  aComputer := WindowsComputer{}
  aComputer.LoadDatabaseInformation(device)

  // Display Output

  fmt.Println(aComputer.DisplayName)

  //fmt.Println(device.DeviceInformation.DisplayName)
}









//Device States
//  1 = No Device Contact
//  2 = Enrollment Began
//  3 = Prestage Enrollment In Progress
//  4 = Initial Configuration/Applications Loading
//  5 = Device Complys To Organisation Security/etc Standards
//  6 = Device Is Enrolled And Healthy
//  7 = Device Is Enrolled And Not Healthy (Lost Communication Or Something)
//  8 = Device Is In Recovery/Lockdown
//  9 = The Device Is In A Bodyless State (The Computers Software is being Reset So The Info Wont Change But The Policys/Apps Will Need to Be Reapplyed)


type Computer struct {
  UUID string
  DisplayName string
  DeviceState int

  //Device Information
  //Network Interfaces
  //Operating System Informaion -> Versions
  //Installed Applications -> Maybe Device Specific
}

func (comp *Computer) LoadDatabaseInformation(_comp Device) {
  comp.DisplayName = _comp.DeviceInformation.DisplayName
}

func (comp *Computer) Checkout() {
  //Unenroll For THe Server Side
  //Maybe Return Error/Nil
}

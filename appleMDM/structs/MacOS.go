package structs

type MacOS_DeviceState struct {
	Token           []byte `sql:"Token,notnull"`
	PushMagic       string `sql:"PushMagic,notnull"`
	UnlockToken     []byte `sql:"PushMagic,notnull"`
	LastUpdate      int64  `sql:"LastUpdate,notnull"`
	EnrollmentState int64  `sql:"EnrollmentState,notnull"`
}

type MacOS_DeviceInfo struct {
	OSVersion    string `sql:"OSVersion,notnull"` //TODO: What Does The notnull Do?
	BuildVersion string `sql:"BuildVersion,notnull"`
	ProductName  string `sql:"ProductName,notnull"`
	SerialNumber string `sql:"SerialNumber,notnull"`
	IMEI         string `sql:"IMEI,notnull"`
	MEID         string `sql:"MEID,notnull"`
}

type MacOS_DeviceConfiguration struct {
	//Profiles     []ProfileListItem `sql:"Profiles,notnull"`
}

//TableName        struct{} `sql:"devices"`
//devices.Computer          // Extend The Default Computer

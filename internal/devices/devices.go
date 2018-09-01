package device

type Computer struct {
	TableName   struct{}    `sql:"devices"`
	UUID        string      `sql:"uuid,pk"`
	DeviceState interface{} `sql:"DeviceState"` //TODO: Maybe Add ",notnull" To These
	DeviceInfo  interface{} `sql:"DeviceInfo"`
}

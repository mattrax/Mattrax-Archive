package device

type Computer struct {
	TableName struct{} `sql:"devices"`
	UUID      string   `sql:"uuid,pk"`
	Name      string
}

func (c Computer) Test() string {
	return "Hello World"
}

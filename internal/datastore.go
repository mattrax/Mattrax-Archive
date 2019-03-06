package mattrax

type DataStore interface {
	SaveDevice(device Device) error
	RetrieveDevice(device Device) error
}

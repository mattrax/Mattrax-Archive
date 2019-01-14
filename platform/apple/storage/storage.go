package storage

import (
	"time"
)

type Service interface {
	CreateDevice(Device) error
	UpdateTokens(uuid string, token []byte, pushMagic string, unlockToken []byte) error
	FindDevice(uuid string) (Device, error)
	DeleteDevice(uuid string) error
	// TODO: close() -> Closes the Postgres Connection
}

type Device struct {
	UUID         string // UUIDv4 + Primary Key
	State        int    // NOT NULL
	Topic        string
	OSVersion    string `db:"os_version"`
	BuildVersion string `db:"build_version"`
	ProductName  string `db:"product_name"`
	SerialNumber string `db:"serial_number"` // NOT NULL
	IMEI         string
	MEID         string
	DeviceName   string `db:"device_name"`
	Challenge    []byte
	Model        string
	ModelName    string    `db:"model_name"`
	CreatedBy    string    `db:"created_by"` // UUIDv4
	CreatedAt    time.Time `db:"created_at"`
	Token        []byte
	PushMagic    string `db:"push_magic"`
	UnlockToken  []byte `db:"unlock_token"`
}

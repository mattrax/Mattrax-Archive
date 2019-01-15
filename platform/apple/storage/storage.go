package storage

import (
	"errors"
	"time"

	"github.com/mattrax/Mattrax/platform/apple/endpoints/structs"
)

type Service interface {
	CreateDevice(Device) error
	UpdateDevice(Device) error
	FindDevice(uuid string) (Device, error)
	DeleteDevice(uuid string) error
	// TODO: close() -> Closes the Postgres Connection
}

type Device struct {
	UUID           string // UUIDv4 + Primary Key
	Enrolled       bool   // NOT NULL
	AwaitingConfig bool   `db:"awaiting_config"` // NOT NULL
	Topic          string
	OSVersion      string `db:"os_version"`
	BuildVersion   string `db:"build_version"`
	ProductName    string `db:"product_name"`
	SerialNumber   string `db:"serial_number"` // NOT NULL
	IMEI           string
	MEID           string
	DeviceName     string `db:"device_name"`
	Challenge      []byte
	Model          string
	ModelName      string    `db:"model_name"`
	CreatedBy      string    `db:"created_by"` // UUIDv4
	CreatedAt      time.Time `db:"created_at"`
	Token          []byte
	PushMagic      string `db:"push_magic"`
	UnlockToken    []byte `db:"unlock_token"`
}

func NewDeviceFromCheckinAuthentication(cmd structs.CheckinCommand) (*Device, error) {
	if cmd.MessageType != "Authenticate" {
		return nil, errors.New("Error can't create a new device from a non 'Authenticate' Checkin Command")
	}

	return &Device{
		UUID:           cmd.UDID,
		Enrolled:       false,
		AwaitingConfig: cmd.AwaitingConfiguration,
		Topic:          cmd.Topic,
		OSVersion:      cmd.OSVersion,
		BuildVersion:   cmd.BuildVersion,
		ProductName:    cmd.ProductName,
		SerialNumber:   cmd.SerialNumber,
		IMEI:           cmd.IMEI,
		MEID:           cmd.MEID,
		DeviceName:     cmd.DeviceName,
		Challenge:      cmd.Challenge, // TOOD: Should This Be Put Into The database?
		Model:          cmd.Model,
		ModelName:      cmd.ModelName,
		CreatedBy:      "00000000-0000-0000-0000-000000000000", // TODO: Implemented This
		CreatedAt:      time.Now(),
	}, nil
}

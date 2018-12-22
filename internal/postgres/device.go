package postgres

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/mattrax/Mattrax/internal/mattrax"
)

type deviceRepository struct {
	db *sql.DB
}

func (r deviceRepository) Create(device *mattrax.Device) error {
	//TODO: Test These Null Checks
	if device.ID == mattrax.DeviceID("") || device.Platform == mattrax.Platform("") || device.PlatformID == "" || device.OSVersion == "" || device.OSEdition == "" || device.AssignedTo == mattrax.UserID("") || device.EnrollmentTime.IsZero() || device.EnrolledBy == mattrax.UserID("") || device.LatestUpdate.IsZero() {
		return mattrax.ErrInvalidDeviceValues
	}

	_, err := r.db.Exec("INSERT INTO devices(id, platform, platform_id, platform_data, OSVersion, OSEdition, DeviceName, SerialNumber, Policies, Applications, AssignedTo, LatestUpdate, EnrolledBy, EnrollmentTime) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);", device.ID, device.Platform, device.PlatformID, device.PlatformData, device.OSVersion, device.OSEdition, device.DeviceName, device.SerialNumber, pq.Array(device.Policies), pq.Array(device.Applications), device.AssignedTo, device.LatestUpdate, device.EnrolledBy, device.EnrollmentTime)

	return err
}

func (r deviceRepository) Update(device *mattrax.Device) error {

	return nil
}

func (r deviceRepository) Remove(device *mattrax.Device) error {
	return nil

}

func (r deviceRepository) Find(id mattrax.DeviceID) (mattrax.Device, error) {

	return mattrax.Device{}, nil
}

func (r deviceRepository) FindAll() ([]mattrax.Device, error) {

	return nil, nil
}

func NewDeviceRepository(db *sql.DB) mattrax.DeviceRepository {
	// TODO: Init Database Schema

	return deviceRepository{db}
}

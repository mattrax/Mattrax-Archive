package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// The postgres SQL driver
	_ "github.com/lib/pq"
	"github.com/mattrax/Mattrax/platform/apple/storage"
)

const schema = `
CREATE TABLE IF NOT EXISTS apple_devices (
	uuid text PRIMARY KEY,
	state int NOT NULL,
	topic text,
	os_version text,
	build_version text,
	product_name text,
	serial_number text NOT NULL UNIQUE,
	imei text UNIQUE,
	meid text UNIQUE,
	device_name text,
	challenge text,
	model text,
	model_name text,
	created_by text,
	created_at time,
	token bytea UNIQUE,
	push_magic text NOT NULL DEFAULT ''::text UNIQUE,
	unlock_token bytea UNIQUE
);` // TODO: created_by use custom foreign key

// TOOD: Warning For Insecure SSLModes
// TODO: Close The Database Connection On Quit

// TODO: dbName not working
func New(host, user, password, dbName, sslMode string) (storage.Service, error) { // TODO: Wrap The Errors + Go Doc + Parse in Port
	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbName, sslMode)
	db, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	// TODO: Do I need to ping the db

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return &postgres{db}, nil
}

type postgres struct {
	db *sqlx.DB
}

func (p *postgres) CreateDevice(device storage.Device) error {
	_, err := p.db.NamedExec(`INSERT INTO apple_devices(uuid, state, topic, os_version, build_version, product_name, serial_number, imei, meid, device_name, challenge, model, model_name, created_by, created_at) VALUES(:uuid, :state, :topic, :os_version, :build_version, :product_name, :serial_number, :imei, :meid, :device_name, :challenge, :model, :model_name, :created_by, :created_at)`, device)
	return err
}

func (p *postgres) UpdateDevice(uuid string, token []byte, pushMagic string, unlockToken []byte) error {
	_, err := p.db.Exec(`UPDATE apple_devices SET token=$2, push_magic=$3, unlock_token=$4 WHERE uuid=$1`, uuid, token, pushMagic, unlockToken)
	return err
}

func (p *postgres) UpdateTokens(uuid string, token []byte, pushMagic string, unlockToken []byte) error {
	_, err := p.db.Exec(`UPDATE apple_devices SET token=$2, push_magic=$3, unlock_token=$4 WHERE uuid=$1`, uuid, token, pushMagic, unlockToken)
	return err
}

// TODO: Touch Minimum Columns in the DB For each transaction

func (p *postgres) FindDevice(uuid string) (storage.Device, error) {
	var device storage.Device
	return device, p.db.Get(&device, "SELECT * FROM apple_devices WHERE uuid=$1", uuid)
	/*var device storage.Device
	err := p.db.QueryRow("SELECT uuid, name FROM apple_devices where uuid=$1 limit 1", uuid).Scan(&device.UUID, &device.Name)
	return &device, err*/ // TODO: Confirm A Blank Devices Doesn't Crash -> Test
}

func (p *postgres) DeleteDevice(uuid string) error {
	return nil
	//_, err := p.db.NamedExec(`INSERT INTO apple_devices(uuid, state, topic, os_version, build_version, product_name, serial_number, imei, meid, device_name, challenge, model, model_name, created_by, created_at) VALUES(:uuid, :state, :topic, :osversion, :buildversion, :productname, :serialnumber, :imei, :meid, :devicename, :challenge, :model, :modelname, :createdby, :createdat)`, device)
	//return err
}

package builtin

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
	mattrax "github.com/mattrax/Mattrax/internal"
)

type dataStore struct {
	db *bolt.DB
}

var (
	DevicesBucket = []byte("devices")
	UsersBucket   = []byte("users")
)

func (ds *dataStore) SaveDevice(device mattrax.Device) error {
	rawDevice, err := device.Serialise()
	if err != nil {
		return err // TODO: Wrap The Error
	}

	err = ds.db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(DevicesBucket).Put(device.ID(), rawDevice)
		return err
	})
	if err != nil {
		return err // TODO: Wrap The Error
	}

	return nil
}

func (ds *dataStore) RetrieveDevice(device mattrax.Device) error {
	// TODO
	return nil
}

func (ds *dataStore) SaveUser(device mattrax.User) error {
	// TODO
	return nil
}

func (ds *dataStore) RetrieveUser(device mattrax.User) error {
	// TODO
	return nil
}

func NewDataStore(path string) mattrax.DataStore {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatal(err) // TODO: Deal With Error
	}

	// TODO: Detect And Setup The Database 'Tables' + Migrations + Add and Check Schema Version
	_ = db.Update(func(tx *bolt.Tx) error { // TODO: Error Handling
		_, err := tx.CreateBucketIfNotExists(DevicesBucket)
		return err
	})

	//defer db.Close() // TODO: Cleanup
	return &dataStore{db}
}

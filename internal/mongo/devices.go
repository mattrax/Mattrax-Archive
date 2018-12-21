package mongo

import (
	"github.com/mattrax/Mattrax/internal/mattrax"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type deviceRepository struct {
	db      string
	session *mgo.Session
}

func (r deviceRepository) Create(device *mattrax.Device) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("devices")

	_, _ = c.Upsert(bson.M{"id": device.ID}, bson.M{"$set": device}) // TODO: Error Handling (second Output valueß)

	return nil
}

func (r deviceRepository) Remove(device *mattrax.Device) error {
	return nil

}

func (r deviceRepository) Find(id mattrax.DeviceID) (*mattrax.Device, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("devices")

	var result mattrax.Device
	if err := c.Find(bson.M{"id": id}).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, mattrax.ErrUnknownDevice
		}
		return nil, err
	}

	return &result, nil
}

func NewDeviceRepository(db string, session *mgo.Session) (mattrax.DeviceRepository, error) {
	r := deviceRepository{db, session}

	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("devices")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}

	return r, nil
}

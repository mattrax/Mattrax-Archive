package apple

import (
	"bytes"
	"encoding/gob"
)

type Mac struct {
	UUID string
	// TODO: The Rest
}

func (d Mac) ID() []byte {
	return []byte("applemac." + d.UUID)
}

func (d Mac) Serialise() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(d); err != nil {
		return nil, err // TODO: Wrap Error
	}
	return buf.Bytes(), nil
}

// TODO: IOS Support In The Future

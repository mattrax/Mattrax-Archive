package windows

import (
	"bytes"
	"encoding/gob"
)

type Windows struct {
	UUID string
	// TODO: The Rest
}

func (d Windows) ID() []byte {
	return []byte("windows." + d.UUID)
}

func (d Windows) Serialise() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(d); err != nil {
		return nil, err // TODO: Wrap Error
	}
	return buf.Bytes(), nil
}

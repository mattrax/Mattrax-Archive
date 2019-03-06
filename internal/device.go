package mattrax

type Device interface {
	ID() (id []byte)
	Serialise() (rawDevice []byte, err error)
}

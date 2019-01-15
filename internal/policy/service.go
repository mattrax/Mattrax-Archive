package policyservice

// type Policy interface {}

type Policy struct {
	UUID       string
	Identifier string
	Raw        []byte
}

func (p *Policy) QueueToDevice(uuid string) error {
	// TODO: Get The Device and Its OS

	return nil
}

// MarshelToApple
// MarshelToWindows

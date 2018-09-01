package appleMDM

import (
	devices "../internal/devices"
)

type MacOS struct {
	devices.Computer // Extend The Default Computer
	PushToken        string
}

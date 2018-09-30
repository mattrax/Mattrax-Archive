package models

// CheckinRequest represents an MDM checkin command struct. /?TODO: Move To The Checkin Handler Package
type CheckinRequest struct {
	MessageType string // Could Be Authenticate or TokenUpdate or CheckOut
	Topic       string
	UDID        string
	// Other Information Is Directly In The Device Struct So It Can Go: Plist In -> Parsed Struct -> Save To DB
}

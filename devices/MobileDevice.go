//TODO Header

package devices

type MobileDevice struct {
  Name string
}

func (r *MobileDevice) Testing() string {
    return "Hello World, Testing"
}

func (r *MobileDevice) InstantAction(action int) string {
  // Use Switch For actions
  return "Hello World, Testing"

  //Actions For Example:
  //  Ping Device (APNS)
  //  Shutdown
  //  Restart
  //  Lock
}


//Send Update (Notification)
//Get Details
//Parse To Postgres Database
//IP Address/Mac Address
//Installed Applications
//Owner/User/Local Users

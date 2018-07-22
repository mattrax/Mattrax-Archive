//TODO Header

package devices

type Computer struct {
  UUID string
  Name string
}

func (r *Computer) Testing() string {
    return "Hello World, Testing"
}

func (r *Computer) InstantAction(action int) string {
  // Use Switch For actions
  return "Hello World, Testing"

  //Actions For Example:
  //  Ping Device (APNS)
  //  Shutdown
  //  Restart
  //  Lock
}
/*
func (r *Computer) DeployPolicy(policy interface) error {
  return nil
}
*/


//Send Update (Notification)
//Get Details
//Parse To Postgres Database
//Network Interfaces/Mac Address
//Installed Applications
//Owner/User/Local Users

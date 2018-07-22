package main

type Device struct {
  Name string
}

func (r *Device) Testing() string {
    return "Hello World, Testing"
}

func (r *Device) InstantAction(action int) string {
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

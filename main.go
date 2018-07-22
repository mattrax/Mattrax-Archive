package main

import (
  "github.com/mattrax/Mattrax/internal" // Mattrax Internal (Logging, Database and Config)
  //"github.com/mattrax/Mattrax/demoMDM"
)

var config, log, pgdb = internal.GetInternalState()

func main() {


  internal.CleanInternalState() //Run On Exit
}

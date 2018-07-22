//This Folder Is A Demo of A Device Class For Both Windows And Apple.

package main

import (
  "fmt"
)

func main() {
  oscarsPhone := Device{ Name: "Testing" }

  fmt.Println(oscarsPhone)
  fmt.Println(oscarsPhone.Name)
  fmt.Println(oscarsPhone.Testing())
}

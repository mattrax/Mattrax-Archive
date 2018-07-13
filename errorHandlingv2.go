package main

import (
  "fmt"
  "log"
  "errors"

  "time"
  "net/http"
  "github.com/gorilla/mux" // HTTP Router
  //"github.com/Benchkram/errz"
)

func main() {
  r := mux.NewRouter()
  r.Handle("/", appHandler(helloWorld))
  //r.HandleFunc("/broken", appHandler(errorCausing))

  srv := &http.Server{
      Handler:      r,
      Addr:         "127.0.0.1:8000",
      // Good practice: enforce timeouts for servers you create!
      WriteTimeout: 15 * time.Second,
      ReadTimeout:  15 * time.Second,
  }

  log.Fatal(srv.ListenAndServe())
}





type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        fmt.Println("HTTPS Error", err.Error())
        http.Error(w, "A Server Side Error Occured", 500)
    }
}





var err error //Having This Globa Will Cause Errors With Simultanous (Web Handlers)


func helloWorld(w http.ResponseWriter, r *http.Request) error {
  /*defer errz.Recover(&err) //recover all panics
  fmt.Println("1")
  err = FailAtSomething()
  fmt.Println("2")
  errz.Fatal(err) //panics on error
  fmt.Println("3")*/

  //return FailAtSomething()

  fmt.Fprintf(w, "Hello World")
  return nil
}

func FailAtSomething() error {
  return errors.New("emit macho dwarf: elf header corrupted")
}

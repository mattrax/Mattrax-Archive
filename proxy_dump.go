package main

import (
  "fmt"
  "io"
  "bytes"
  "io/ioutil"

  "net/http"
  //"net/http/httputil"
)

var (
  listen = ":8000"
  proxy_to = "http://localhost:8001"
  proxy_to_domain = "mdm.otbeaumont.me"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request: " + r.RemoteAddr + " " + r.Method + " " + r.URL.String())

    client := &http.Client {}
    request, err := http.NewRequest(r.Method, proxy_to + r.URL.String(), nil)
    request.Host = proxy_to_domain

    resp, err2 := client.Do(request)
    if err != nil { fmt.Println("Error Doing Get", err); http.Error(w, "", http.StatusServiceUnavailable); return }
    if err2 != nil { fmt.Println("Error Doing Get", err2); http.Error(w, "", http.StatusServiceUnavailable); return }

    defer resp.Body.Close()
    copyHeader(w.Header(), resp.Header)
    w.WriteHeader(resp.StatusCode)
    b := bytes.NewBuffer(make([]byte, 0))
    reader := io.TeeReader(resp.Body, b)
    io.Copy(w, reader)
    resp.Body = ioutil.NopCloser(b)
    body, err := ioutil.ReadAll(resp.Body)

    if bodyStr := string(body); bodyStr != "" { fmt.Println(bodyStr) }
}

func copyHeader(dst, src http.Header) {
    for k, vv := range src {
        for _, v := range vv {
            dst.Add(k, v)
        }
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Proxy Listing At Port: " + listen + " And Redirecting To Port " + proxy_to)
    http.ListenAndServe(listen, nil)
}

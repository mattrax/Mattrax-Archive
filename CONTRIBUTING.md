# Contributing

Welcome to Mattrax! This file contains information that will help you if you are looking to develop on the Mattrax codebase.  If you are working on the project and are having problems please feel free to [contact me (Oscar Beaumont)](https://otbeaumont.me/contact).


## Things that might help you

Here's a few links to get started and to help you understand how the project works.

- Read through the [Apple MDM Protocol Reference](https://developer.apple.com/library/content/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html) for a deeper understanding of how Apple devices are managed.
- Read through the [Windows MDM Reference Reference](https://developer.apple.com/business/documentation/MDM-Protocol-Reference.pdf) for a deeper understanding of how Windows 10 devices are managed.
- If you run into a problem that you're not sure how to fix, browse through the [the issue tracker](https://github.com/mattrax/Mattrax/issues) for anyone else having the same problem and if not create a new issue.
- If you come across something that others might benefit from? Considering updating or writing a [wiki page](https://github.com/mattrax/Mattrax/wiki).

## A Couple Of Design Focus's
* Clean and Fast Code!
* Production Ready - Stable and Predictable
* A Means to Some What Debugging A Failure After It Happens
* Each Server Instance Should Be Stateless
* Follow The Clean Architecture. For more information [click here](https://otbeaumont.me/blog/Go-Lang-Clean-Architecture-and-Domain-Driven-Development/).

## Running the project

To run Mattrax from the sources, you will need the latest version of [Go Lang](https://golang.org/dl/) installed. Then run the commands:

```
git clone https://github.com/mattrax/Mattrax.git && cd Mattrax
go run ./cmd/mattrax
```

## Go Resources

A few helpful resources for getting started with Go:

* [Writing, building, installing, and testing Go code](https://www.youtube.com/watch?v=XCsL89YtqCs)
* [Resources for new Go programmers](http://dave.cheney.net/resources-for-new-go-programmers)
* [How I start](https://howistart.org/posts/go/1)
* [How to write Go code](https://golang.org/doc/code.html)

A few helpful resources for the development style of this repository:

* [Building an enterprise service in Go](https://www.youtube.com/watch?v=twcDf_Y2gXY) - [The code the talk is about](https://github.com/marcusolsson/goddd/)
* [How I write Go HTTP services after seven years](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)

## Important libraries and frameworks

Mattrax is built using popular Go packages outside the standard libraries that might be worth checking out before working on the code. They can be found in the [Gopkg.toml](Gopkg.toml) file.

## A handy code snippet
This code prints the body of the http.Request to the console. This is handy for debugging what the MDM enrolled device is sending to the server.

```
bodyBytes, _ := ioutil.ReadAll(r.Body)
log.Println(string(bodyBytes))
log.Println()
return
```

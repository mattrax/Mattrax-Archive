package main

import (
	"fmt"
	"runtime"
)

var (
	version   string
	gitCommit string
	gitState  string
	buildTime string
)

func displayVersion() {
	if version == "" || gitCommit == "" || gitState == "" || buildTime == "" {
		fmt.Println("WARNING: This binary is either damaged or was build during development!")
	}

	fmt.Println(`Mattrax:
    Version:      ` + version + `
    Go version:   ` + runtime.Version() + `
    Git commit:   ` + gitCommit + `
    Git state:    ` + gitState + `
    Built:        ` + buildTime + `
    OS/Arch:      ` + runtime.GOOS + `/` + runtime.GOARCH)
}

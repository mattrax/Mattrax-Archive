package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// TODO: Inject Versions + Git Information at build time

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "help", "-help", "--help":
		help()
	case "version", "-version", "--version":
		version()
	case "server":
		server(os.Args[2:]) //TODO: Error Handling
	default:
		usage()
		os.Exit(1)
	}
}

func help() { //TODO
	helpText := `USAGE: mattrax <COMMAND>
Available Commands:
	server
	version
Use mattrax <command> -h for additional usage of each command.
Example: mattrax serve -h
`
	fmt.Println(helpText)
}

func version() { //TODO
	helpText := `USAGE: mattrax <COMMAND>
Available Commands:
	server
	version
Use mattrax <command> -h for additional usage of each command.
Example: mattrax serve -h
`
	fmt.Println(helpText)
}

func usage() { //TODO
	helpText := `USAGE: mattrax <COMMAND>
Available Commands:
	server
	version
Use mattrax <command> -h for additional usage of each command.
Example: mattrax serve -h
`
	fmt.Println(helpText)
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}

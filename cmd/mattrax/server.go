package main

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func server(args []string) error {
	// Load The Configuration
	flagset := flag.NewFlagSet("server", flag.ExitOnError)
	var (
	/*flConfigPath    = flagset.String("config-path", "/var/db/mattrax", "path to configuration directory")
	flServerURL     = flagset.String("server-url", "", "public HTTPS url of your server")
	flTLSCert       = flagset.String("tls-cert", "", "path to TLS certificate")
	flTLSKey        = flagset.String("tls-key", "", "path to TLS private key")
	flAddr          = flagset.String("http-addr", ":8000", "https listen address of mdm server. defaults to :8000")
	flHomePageRedir = flagset.Bool("homepage-redir", true, "redirect request to the home page to an alternate url")*/
	//flDebug = flagset.Bool("debug", false, "enables debug mode which log extra data and enabled features designed to aid development")
	)
	flagset.Usage = usageFor(flagset, "mattrax server [flags]")
	if err := flagset.Parse(args); err != nil {
		return err
	}

	///// From here down is cleaned

	// Setup The Logging
	/*zerolog.TimeFieldFormat = ""

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *flDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})*/

	//output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	/*output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}*/
	//output.FormatMessage = func(i interface{}) string {
	//	return fmt.Sprintf("***%s****", i)
	//}
	/*output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}*/

	console := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	file := zerolog.New(os.Stderr) //.With().Timestamp().Logger()

	logger := zerolog.New(io.MultiWriter(file, console)).With().Timestamp().Caller().Logger()

	// BREAK

	/*

		sublogger := log.With().
		                 Str("component", "foo").
		                 Logger()
	*/

	logger.Info().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	// if e := log.Debug(); e.Enabled() {

	return nil
}

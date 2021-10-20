package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/jylitalo/rotakey"
)

func main() {
	flag.Bool("debug", false, "Debug output")
	flag.BoolP("verbose", "v", false, "Verbose output")
	flag.Parse()
	debug, errD := flag.CommandLine.GetBool("debug")
	verbose, errV := flag.CommandLine.GetBool("verbose")
	if err := coalesceError(errD, errV); err != nil {
		flag.PrintDefaults()
		// PrintDefaults does exit
	}
	switch {
	case debug:
		log.SetLevel(log.DebugLevel)
	case verbose:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
	if err := rotakey.Execute(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task completed.")
}

func coalesceError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

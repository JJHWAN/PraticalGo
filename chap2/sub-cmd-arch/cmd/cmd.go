package cmd

import (
	"flag"
	"fmt"
	"io"
)

type httpConfig struct {
	url  string
	verb string
}

func HandleHttp(w io.Writer, args []string) error {
	var v string

	// flag.ContinueOnError
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "GET", "HTTP method")

	fs.Usage = func() {
		var usageString = `
http: A HTTP Client

http: <options> server`
		fmt.Fprint(w, usageString)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	err := fs.Parse(args)
	if err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return ErrorNoServerSpecified
	}

	c := httpConfig{verb: v}
	c.url = fs.Arg(0)
	fmt.Fprintln(w, "Executing http command")
	return nil

}

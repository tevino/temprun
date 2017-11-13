package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/tevino/temprun/command"
	"github.com/tevino/temprun/template"
)

type Args struct {
	Conf    string
	Prefix  string
	Src     string
	Dst     string
	DstMode uint
	Cmd     []string
}

const usageSSSS = `Usage of %s:

%s [flags] [[command] [argument ...]]

Description:
  %s renders template with environment variables into destination, then optionally execute command with its following arguments.

Template:
  If configuration file is not specified, template will be read from standard input.

Destination:
  destination is set to standard output by default.

Command:
  Command will be executed with the same environment variables as %s itself.

Flags:
`

func parseArgs() *Args {
	args := &Args{}
	flagset := flag.NewFlagSet("temprun", flag.ExitOnError)
	flagset.StringVar(&args.Conf, "conf", "temprun.conf", "Configuration file")
	flagset.StringVar(&args.Prefix, "prefix", "", "Environment prefix")
	flagset.StringVar(&args.Src, "src", "-", "Template file")
	flagset.StringVar(&args.Dst, "dst", "", "Destination file")
	flagset.UintVar(&args.DstMode, "dst-mode", 0644, "File mode of destination file")
	flagset.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, usageSSSS, name, name, name, name)
		flagset.PrintDefaults()
	}
	flagset.Parse(os.Args[1:])
	args.Cmd = flagset.Args()
	return args
}

func validateArgs(args *Args) error {
	return nil
}

func openDst(dst string, dstMode uint) (io.WriteCloser, error) {
	var dstWriter io.WriteCloser
	dstWriter = os.Stdout
	if dst != "" {
		dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(dstMode))
		if err != nil {
			return nil, err
		}
		dstWriter = dstFile
		defer dstFile.Close()
	}
	return dstWriter, nil
}

func openSrc(src string) (io.ReadCloser, error) {
	if src == "-" {
		return os.Stdin, nil
	}
	return os.Open(src)
}

func renderTemplate(args *Args) error {
	if fileExists(args.Conf) {
		// use configuration
		return errors.New("Configuration is not implemented yet")
	}

	src, err := openSrc(args.Src)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := openDst(args.Dst, args.DstMode)
	if err != nil {
		return err
	}
	defer dst.Close()

	tpl := template.NewEnvTemplate(os.Environ(), args.Prefix, nil)
	return tpl.RenderToWriter(src, dst)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func main() {
	args := parseArgs()
	if err := validateArgs(args); err != nil {
		log.Fatal("Invalid arguments: ", err)
	}
	if err := renderTemplate(args); err != nil {
		log.Fatal("Error rendering template: ", err)
	}
	if err := command.ExecCmd(args.Cmd); err != nil {
		log.Fatal("Error executing command: ", err)
	}
}

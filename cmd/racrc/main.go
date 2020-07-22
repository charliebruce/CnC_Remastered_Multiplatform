package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func printRootUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), `usage: %s <command>
Available commands:
  %s - Calculate a CRC for RA
`, path.Base(os.Args[0]), calcFs.Name())
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printRootUsage
	flag.Parse()
	rootCommand := ""
	if flag.NArg() > 0 {
		rootCommand = flag.Arg(0)
	}

	if rootCommand == calcFs.Name() {
		calcCommand(flag.Args()[1:])
		return
	}

	flag.Usage()
	os.Exit(2)
}

package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/crc"
)

var (
	calcFs    *flag.FlagSet
	formatHex *bool
	formatInt *bool
)

func init() {
	calcFs = flag.NewFlagSet("calc", flag.ContinueOnError)
	formatHex = calcFs.Bool("X", true, "Format output to hex uppercase")
	formatInt = calcFs.Bool("I", false, "Format output to int32")
}

func printCalcUsage() {
	fmt.Fprintf(calcFs.Output(), `usage: %s calc [<option>, ...] <string>
Options:
`, path.Base(os.Args[0]))
	calcFs.PrintDefaults()
}

func calcCommand(args []string) {
	calcFs.Usage = printCalcUsage
	calcFs.Parse(args)

	if calcFs.NArg() != 1 || calcFs.Arg(0) == "" {
		calcFs.Usage()
		os.Exit(1)
	}

	s := calcFs.Arg(0)

	if *formatInt {
		println(crc.RA(s))
	} else {
		println(crc.RAHexString(s))
	}
}

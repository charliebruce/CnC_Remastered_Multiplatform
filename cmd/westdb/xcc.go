package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/db"
	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/file"
	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/ops"
	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/xcc"
)

var (
	xccFlagSet = flag.NewFlagSet("xcc", flag.ContinueOnError)
)

func printXccUsage() {
	fmt.Fprintf(xccFlagSet.Output(), `usage: %s xcc <command>
Available commands:
  import - Imports from XCC Database (and save to disk)
`, path.Base(os.Args[0]))
	xccFlagSet.PrintDefaults()
}

func xccCommand(args []string) {
	xccFlagSet.Usage = printXccUsage

	if len(args) != 1 || args[0] != "import" {
		xccFlagSet.Usage()
		os.Exit(1)
	}

	var xdb *xcc.Database
	var fe *file.Entries

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		r := xcc.MustGetFromRemote()
		defer r.Close()
		xdb = xcc.MustReadDatabase(r)
	}()

	go func() {
		defer wg.Done()
		fe = mustReadEntries()
	}()

	wg.Wait()
	ops.MustMerge(xdb, fe)
	mustWriteToYaml(fe)
}

func mustReadEntries() *file.Entries {
	f, err := os.Open("files.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			return file.NewEntries([]file.Entry{})
		}
	}
	defer f.Close()

	r := bufio.NewReader(f)
	return db.MustReadYAMLDatabase(r)
}

func mustWriteToYaml(fe *file.Entries) {
	file, err := os.OpenFile("files.yaml", os.O_CREATE|os.O_WRONLY, 0611)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer func() {
		err := w.Flush()
		if err != nil {
			log.Println(err)
		}
	}()

	db.WriteEntriestoYAML(fe, w)
}

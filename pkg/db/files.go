package db

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"unsafe"

	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/file"
	"gopkg.in/yaml.v2"
)

type yamlDatabaseV1 struct {
	APIVersion  string          `yaml:"apiVersion"`
	Kind        string          `yaml:"kind"`
	FileEntries []YAMLFileEntry `yaml:"files"`
}

type YAMLFileEntry struct {
	Filename    string
	EmbeddedIn  file.ShaDigest
	ContentHash file.ShaDigest
	Comment     string
	RaCRC       uint32 // Only unsigned can converted from YAML
}

func MustReadYAMLDatabase(r io.Reader) *file.Entries {

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}

	db := yamlDatabaseV1{}

	err = yaml.Unmarshal(buf, &db)
	if err != nil {
		log.Fatal(err)
	}

	if db.Kind != "Files" || db.APIVersion != "1" {
		if len(db.FileEntries) == 0 {
			return file.NewEntries([]file.Entry{})
		}
		log.Fatalf("Unsupported %s %s", db.Kind, db.APIVersion)
	}

	log.Printf("Read %d file entries\n", len(db.FileEntries))

	return file.NewEntries(convertEntries(db.FileEntries))
}

func convertEntries(yamlEntries []YAMLFileEntry) []file.Entry {
	fes := make([]file.Entry, len(yamlEntries), len(yamlEntries))
	for i, ye := range yamlEntries {
		ip := unsafe.Pointer(&ye.RaCRC)
		fes[i] = *file.NewEntry(ye.Filename, *(*int32)(ip))
		fes[i].Comment = ye.Comment
		fes[i].ContentHash = ye.ContentHash
		fes[i].EmbeddedIn = ye.EmbeddedIn
	}

	return fes
}

func writeHeader(w io.Writer) {
	fmt.Fprintln(w, `apiVersion: 1
kind: Files`)
}

func writeEntry(entry *file.Entry, w io.Writer) {
	fmt.Fprintf(w, `- racrc: 0x%08X
  filename: "%s"
`, entry.Uint32CRC(), entry.Filename())

	if !entry.ContentHash.IsZero() {
		fmt.Fprintf(w, `  contentHash: %s
`, entry.ContentHash)
	}

	if !entry.EmbeddedIn.IsZero() {
		fmt.Fprintf(w, `  embeddedIn: %s
`, entry.EmbeddedIn)
	}

	if entry.Comment != "" {
		fmt.Fprintf(w, `  comment: "%s"
`, entry.Comment)
	}
}

func WriteEntriestoYAML(entries *file.Entries, w io.Writer) {

	writeHeader(w)

	fmt.Fprintln(w, "files:")

	entries.Update()

	writeSortedByCRC(entries.Unknown, w)
	log.Printf("Wrote %d file entries without sha\n", len(entries.Unknown))

	writeSortedByContentHash(entries.Known, w)
	log.Printf("Wrote %d file entries with sha\n", len(entries.Known))
}

func writeSortedByCRC(m map[int32]*file.Entry, w io.Writer) {
	keys := make([]int32, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(file.CRCs(keys))

	for _, key := range keys {
		entry := m[key]
		writeEntry(entry, w)
	}
}

func writeSortedByContentHash(m map[file.ShaDigest]*file.Entry, w io.Writer) {
	keys := make([]file.ShaDigest, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(file.ShaDigests(keys))

	for _, key := range keys {
		entry := m[key]
		writeEntry(entry, w)
	}
}

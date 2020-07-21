package file

import (
	"crypto/sha256"
	"log"
	"unsafe"
)

type Entry struct {
	EmbeddedIn  ShaDigest
	ContentHash ShaDigest
	Comment     string
	filename    string
	raCrc       int32
	hasCrc      bool
}

type Entries struct {
	Entries []Entry
	Known   map[ShaDigest]*Entry
	Unknown map[int32]*Entry
}

type ShaDigest [sha256.Size]byte

func NewEntry(filename string, racrc int32) *Entry {
	return &Entry{
		filename: filename,
		raCrc:    racrc,
		hasCrc:   true,
	}
}

func (e *Entry) Filename() string {
	return e.filename
}

func (e *Entry) Int32CRC() int32 {
	return e.raCrc
}

func (e *Entry) Uint32CRC() uint32 {
	ip := unsafe.Pointer(&e.raCrc)
	return *(*uint32)(ip)
}

func (e *Entry) HasCRC() bool {
	return e.hasCrc
}

func (d ShaDigest) IsZero() bool {
	for i := 0; i != sha256.Size; i++ {
		if d[i] != 0 {
			return false
		}
	}
	return true
}

func NewEntries(entries []Entry) *Entries {
	res := &Entries{
		Entries: entries,
		Known:   map[ShaDigest]*Entry{},
		Unknown: map[int32]*Entry{},
	}
	res.Update()
	return res
}

type CRCs []int32

func (c CRCs) Len() int {
	return len(c)
}

func (c CRCs) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c CRCs) Swap(i, j int) {
	v := c[i]
	c[i] = c[j]
	c[j] = v
}

type ShaDigests []ShaDigest

func (d ShaDigests) Len() int {
	return len(d)
}

func (d ShaDigests) Less(i, j int) bool {
	id := d[i]
	jd := d[j]
	for index := 0; index != sha256.Size; index++ {
		if id[index] < jd[index] {
			return true
		}

		if id[index] == jd[index] {
			continue
		}

		return false
	}
	return false
}

func (d ShaDigests) Swap(i, j int) {
	v := d[i]
	d[i] = d[j]
	d[j] = v
}

func (e *Entries) Update() {

	// Check whether all entries are valid
	for _, entry := range e.Entries {
		if !entry.HasCRC() {
			log.Fatalf("Entry for %s has no CRC!\n", entry.Filename())
		}
	}

	for k := range e.Unknown {
		delete(e.Unknown, k)
	}

	// Currently don't see a reason to ever lose Known

	for _, entry := range e.Entries {
		le := entry
		if entry.ContentHash.IsZero() {
			e.Unknown[entry.Int32CRC()] = &le
		} else {
			e.Known[entry.ContentHash] = &le
		}
	}
}

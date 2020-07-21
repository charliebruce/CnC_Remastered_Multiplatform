package ops

// #include <stdlib.h>
// #include <string.h>
// #include "../../REDALERT/platformlib/PLATCRC.H"
import "C"

import (
	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/crc"
	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/file"
	"github.com/abergmeier/CnC_Remastered_Multiplatform/pkg/xcc"
)

func MustMerge(xdb *xcc.Database, y *file.Entries) {

	m := map[int32]int{}

	for i, ye := range y.Entries {
		m[ye.Int32CRC()] = i
	}

	for _, xe := range xdb.Entries {
		c := crc.RA(string(xe.Filename))
		i, ok := m[c]
		if ok {
			y.Entries[i].Comment = string(xe.Comment)
		} else {
			n := file.NewEntry(string(xe.Filename), c)
			n.Comment = string(xe.Comment)
			y.Entries = append(y.Entries, *n)
			i = len(y.Entries) - 1
			m[c] = i
		}

		// Paranoia testing
		if string(xe.Filename) != y.Entries[i].Filename() {
			panic("Filename mismatch")
		}
	}
}

package crc

import (
	"testing"
)

func TestRaHex(t *testing.T) {
	hs := RAHexString("boaticon.shp")
	if hs != "0x3DEC16C9" {
		t.Fatalf("Calculated wrong CRC for boaticon.shp: %s", hs)
	}
}

func TestRa(t *testing.T) {
	i := RA("bigcheck.shp")
	if i != -202636601 {
		t.Fatalf("Calculated wrong CRC for bigcheck.shp: %d (expected: -202636601)", i)
	}
}

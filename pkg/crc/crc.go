package crc

// #include <stdlib.h>
// #include <string.h>
// #include "../../REDALERT/platformlib/PLATCRC.H"
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

func RA(s string) int32 {
	us := strings.ToUpper(s)
	cs := C.CString(us)
	defer C.free(unsafe.Pointer(cs))
	return int32(C.Calculate_CRC(unsafe.Pointer(cs), C.long(C.strlen(cs))))
}

func RAHexString(s string) string {
	crc := RA(s)
	ip := unsafe.Pointer(&crc)
	return fmt.Sprintf("0x%08X", *(*uint32)(ip))
}

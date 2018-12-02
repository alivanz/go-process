// +build !windows

package process

/*
#include <stdlib.h>
#include <stdio.h>


*/
import "C"
import (
	"os"
	"unsafe"
)

// Popen is a system call for popen
func Popen(command string, mode string) *os.File {
	ccommand := C.CString(command)
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(ccommand))
	defer C.free(unsafe.Pointer(cmode))
	cfile := C.popen(ccommand, cmode)
	fd := C.fileno(cfile)
	return os.NewFile(uintptr(fd), "popen file")
}

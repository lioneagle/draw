package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libuser32 = syscall.NewLazyDLL("user32.dll")
	//libmsimg32 = syscall.NewLazyDLL("msimg32.dll")

	// Functions
	procFillRect = libuser32.NewProc("FillRect")
)

// https://msdn.microsoft.com/en-us/library/windows/desktop/dd162719(v=vs.85).aspx
func FillRect(hdc HDC, rect *RECT, brush HBRUSH) int32 {
	ret, _, _ := procFillRect.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(rect)),
		uintptr(brush))

	if ret == 0 {
		panic("FillRect failed")
	}

	return int32(ret)
}

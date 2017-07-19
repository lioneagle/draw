package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libgdi32 = syscall.NewLazyDLL("gdi32.dll")
	//libmsimg32 = syscall.NewLazyDLL("msimg32.dll")

	// Functions
	procCreateDC               = libgdi32.NewProc("CreateDCW")
	procCreateCompatibleDC     = libgdi32.NewProc("CreateCompatibleDC")
	procDeleteDC               = libgdi32.NewProc("DeleteDC")
	procGetObject              = libgdi32.NewProc("GetObjectW")
	procCreateCompatibleBitmap = libgdi32.NewProc("CreateCompatibleBitmap")
)

func CreateDC(lpszDriver, lpszDevice, lpszOutput *uint16, lpInitData *DEVMODE) HDC {
	ret, _, _ := procCreateDC.Call(
		uintptr(unsafe.Pointer(lpszDriver)),
		uintptr(unsafe.Pointer(lpszDevice)),
		uintptr(unsafe.Pointer(lpszOutput)),
		uintptr(unsafe.Pointer(lpInitData)))

	return HDC(ret)
}

func CreateCompatibleDC(hdc HDC) HDC {
	ret, _, _ := procCreateCompatibleDC.Call(
		uintptr(hdc))

	if ret == 0 {
		panic("Create compatible DC failed")
	}

	return HDC(ret)
}

func DeleteDC(hdc HDC) bool {
	ret, _, _ := procDeleteDC.Call(
		uintptr(hdc))

	return ret != 0
}

func GetObject(hgdiobj HGDIOBJ, cbBuffer uintptr, lpvObject unsafe.Pointer) int32 {
	ret, _, _ := procGetObject.Call(
		uintptr(hgdiobj),
		uintptr(cbBuffer),
		uintptr(lpvObject))

	return int32(ret)
}

func CreateCompatibleBitmap(hdc HDC, width, height uint32) HBITMAP {
	ret, _, _ := procCreateCompatibleBitmap.Call(
		uintptr(hdc),
		uintptr(width),
		uintptr(height))

	return HBITMAP(ret)
}

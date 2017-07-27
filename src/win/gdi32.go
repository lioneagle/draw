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
	procCreateDC           = libgdi32.NewProc("CreateDCW")
	procCreateCompatibleDC = libgdi32.NewProc("CreateCompatibleDC")
	procDeleteDC           = libgdi32.NewProc("DeleteDC")

	procGetObject    = libgdi32.NewProc("GetObjectW")
	procSelectObject = libgdi32.NewProc("SelectObject")
	procDeleteObject = libgdi32.NewProc("DeleteObject")

	procCreateCompatibleBitmap = libgdi32.NewProc("CreateCompatibleBitmap")

	procGetStockObject = libgdi32.NewProc("GetStockObject")
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

func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	ret, _, _ := procSelectObject.Call(
		uintptr(hdc),
		uintptr(hgdiobj))

	if ret == 0 {
		panic("SelectObject failed")
	}

	return HGDIOBJ(ret)
}

func DeleteObject(hgdiobj HGDIOBJ) bool {
	ret, _, _ := procDeleteObject.Call(
		uintptr(hgdiobj))

	return ret != 0
}

func CreateCompatibleBitmap(hdc HDC, width, height uint32) HBITMAP {
	ret, _, _ := procCreateCompatibleBitmap.Call(
		uintptr(hdc),
		uintptr(width),
		uintptr(height))

	return HBITMAP(ret)
}

func GetStockObject(fnObject int32) HGDIOBJ {
	ret, _, _ := procGetStockObject.Call(
		uintptr(fnObject))

	return HGDIOBJ(ret)
}

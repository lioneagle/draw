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
	procSelectObject           = libgdi32.NewProc("SelectObject")
	procDeleteObject           = libgdi32.NewProc("DeleteObject")
	procCreateCompatibleBitmap = libgdi32.NewProc("CreateCompatibleBitmap")
	procGetStockObject         = libgdi32.NewProc("GetStockObject")
	procExtCreatePen           = libgdi32.NewProc("ExtCreatePen")
	procCreateBrushIndirect    = libgdi32.NewProc("CreateBrushIndirect")
	procGetDeviceCaps          = libgdi32.NewProc("GetDeviceCaps")
	procSetBkMode              = libgdi32.NewProc("SetBkMode")
	procSetStretchBltMode      = libgdi32.NewProc("SetStretchBltMode")
	procSetBrushOrgEx          = libgdi32.NewProc("SetBrushOrgEx")
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

// https://msdn.microsoft.com/zh-cn/library/windows/desktop/dd144925(d=printer,v=vs.85).aspx
func GetStockObject(fnObject int32) HGDIOBJ {
	ret, _, _ := procGetStockObject.Call(
		uintptr(fnObject))

	return HGDIOBJ(ret)
}

func ExtCreatePen(dwPenStyle, dwWidth uint32, lplb *LOGBRUSH, dwStyleCount uint32, lpStyle *uint32) HPEN {
	ret, _, _ := procExtCreatePen.Call(
		uintptr(dwPenStyle),
		uintptr(dwWidth),
		uintptr(unsafe.Pointer(lplb)),
		uintptr(dwStyleCount),
		uintptr(unsafe.Pointer(lpStyle)))

	return HPEN(ret)
}

func CreateBrushIndirect(lplb *LOGBRUSH) HBRUSH {
	ret, _, _ := procCreateBrushIndirect.Call(
		uintptr(unsafe.Pointer(lplb)))

	return HBRUSH(ret)
}

func GetDeviceCaps(hdc HDC, index int32) int32 {
	ret, _, _ := procGetDeviceCaps.Call(
		uintptr(hdc),
		uintptr(index))

	return int32(ret)
}

func SetBkMode(hdc HDC, iBkMode int32) int32 {
	ret, _, _ := procSetBkMode.Call(
		uintptr(hdc),
		uintptr(iBkMode))

	if ret == 0 {
		panic("SetBkMode failed")
	}

	return int32(ret)
}

func SetStretchBltMode(hdc HDC, iStretchMode int32) int32 {
	ret, _, _ := procSetStretchBltMode.Call(
		uintptr(hdc),
		uintptr(iStretchMode))

	return int32(ret)
}

func SetBrushOrgEx(hdc HDC, nXOrg, nYOrg int32, lppt *POINT) bool {
	ret, _, _ := procSetBrushOrgEx.Call(
		uintptr(hdc),
		uintptr(nXOrg),
		uintptr(nYOrg),
		uintptr(unsafe.Pointer(lppt)))

	return ret != 0
}

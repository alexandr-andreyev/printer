// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package printer

import "unsafe"
import "syscall"

var _ unsafe.Pointer

var (
	modwinspool = syscall.NewLazyDLL("winspool.drv")

	procGetDefaultPrinterW = modwinspool.NewProc("GetDefaultPrinterW")
	procClosePrinter       = modwinspool.NewProc("ClosePrinter")
	procOpenPrinterW       = modwinspool.NewProc("OpenPrinterW")
	procStartDocPrinterW   = modwinspool.NewProc("StartDocPrinterW")
	procEndDocPrinter      = modwinspool.NewProc("EndDocPrinter")
	procWritePrinter       = modwinspool.NewProc("WritePrinter")
	procStartPagePrinter   = modwinspool.NewProc("StartPagePrinter")
	procEndPagePrinter     = modwinspool.NewProc("EndPagePrinter")
	procEnumPrintersW      = modwinspool.NewProc("EnumPrintersW")
)

func GetDefaultPrinter(buf *uint16, bufN *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetDefaultPrinterW.Addr(), 2, uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(bufN)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ClosePrinter(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procClosePrinter.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func OpenPrinter(name *uint16, h *syscall.Handle, defaults uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procOpenPrinterW.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(h)), uintptr(defaults))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func StartDocPrinter(h syscall.Handle, level uint32, docinfo *DOC_INFO_1) (err error) {
	r1, _, e1 := syscall.Syscall(procStartDocPrinterW.Addr(), 3, uintptr(h), uintptr(level), uintptr(unsafe.Pointer(docinfo)))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func EndDocPrinter(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procEndDocPrinter.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WritePrinter(h syscall.Handle, buf *byte, bufN uint32, written *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWritePrinter.Addr(), 4, uintptr(h), uintptr(unsafe.Pointer(buf)), uintptr(bufN), uintptr(unsafe.Pointer(written)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func StartPagePrinter(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procStartPagePrinter.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func EndPagePrinter(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procEndPagePrinter.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func EnumPrinters(flags uint32, name *uint16, level uint32, buf *byte, bufN uint32, needed *uint32, returned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procEnumPrintersW.Addr(), 7, uintptr(flags), uintptr(unsafe.Pointer(name)), uintptr(level), uintptr(unsafe.Pointer(buf)), uintptr(bufN), uintptr(unsafe.Pointer(needed)), uintptr(unsafe.Pointer(returned)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
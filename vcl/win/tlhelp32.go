package win

import "unsafe"

// The th32ProcessID argument is only used if TH32CS_SNAPHEAPLIST or
// TH32CS_SNAPMODULE is specified. th32ProcessID == 0 means the current
// process.
//
// NOTE that all of the snapshots are global except for the heap and module
//  lists which are process specific. To enumerate the heap or module
//  state for all WIN32 processes call with TH32CS_SNAPALL and the
//  current process. Then for each process in the TH32CS_SNAPPROCESS
//  list that isn't the current process, do a call with just
//  TH32CS_SNAPHEAPLIST and/or TH32CS_SNAPMODULE.
//
// dwFlags
//
const (
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPALL      = TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE
	TH32CS_INHERIT      = 0x80000000
)

type TProcessEntry32W struct {
	DwSize              uint32
	CntUsage            uint32
	Th32ProcessID       uint32 // this process
	Th2DefaultHeapID    uintptr
	Th32ModuleID        uint32 // associated exe
	CntThreads          uint32
	Th32ParentProcessID uint32 // this process's parent process
	PcPriClassBase      uint32 // Base priority of process's threads
	DwFlags             uint32
	SzExeFile           [MAX_PATH]uint16 // Path

}

var (
	_CreateToolhelp32Snapshot = kernel32dll.NewProc("CreateToolhelp32Snapshot")
	_Process32FirstW          = kernel32dll.NewProc("Process32FirstW")
	_Process32NextW           = kernel32dll.NewProc("Process32NextW")
)

// CreateToolhelp32SnapShot
func CreateToolhelp32SnapShot(dwFlags, th32ProcessID uint32) uintptr {
	r, _, _ := _CreateToolhelp32Snapshot.Call(uintptr(dwFlags), uintptr(th32ProcessID))
	return r
}

// Process32FirstW
func Process32FirstW(hSnapshot uintptr, lppe *TProcessEntry32W) bool {
	r, _, _ := _Process32FirstW.Call(hSnapshot, uintptr(unsafe.Pointer(lppe)))
	return r != 0
}

// Process32NextW
func Process32NextW(hSnapshot uintptr, lppe *TProcessEntry32W) bool {
	r, _, _ := _Process32NextW.Call(hSnapshot, uintptr(unsafe.Pointer(lppe)))
	return r != 0
}

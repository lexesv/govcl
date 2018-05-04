// +build windows

package win

type TOSVersionInfoExW struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16 // Maintenance UnicodeString for PSS usage
	ServicePackMajor  uint16
	ServicePackMinor  uint16
	SuiteMask         uint16
	ProductType       uint8
	Reserved          uint8
}

type TSystemInfo struct {

	//0: (
	//dwOemId: DWORD);
	//1: (
	ProcessorArchitecture     uint16
	Reserved                  uint16
	PageSize                  uint32
	MinimumApplicationAddress uintptr
	MaximumApplicationAddress uintptr
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

type TVSFixedFileInfo struct {
	Signature        uint32 // e.g. $feef04bd
	StrucVersion     uint32 // e.g. $00000042 = "0.42"
	FileVersionMS    uint32 // e.g. $00030075 = "3.75"
	FileVersionLS    uint32 // e.g. $00000031 = "0.31"
	ProductVersionMS uint32 // e.g. $00030010 = "3.10"
	ProductVersionLS uint32 // e.g. $00000031 = "0.31"
	FileFlagsMask    uint32 // = $3F for version "0.42"
	FileFlags        uint32 // e.g. VFF_DEBUG | VFF_PRERELEASE
	FileOS           uint32 // e.g. VOS_DOS_WINDOWS16
	FileType         uint32 // e.g. VFT_DRIVER
	FileSubtype      uint32 // e.g. VFT2_DRV_KEYBOARD
	FileDateMS       uint32 // e.g. 0
	FileDateLS       uint32 // e.g. 0
}

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

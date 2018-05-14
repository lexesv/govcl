// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (

	// advapi32.dll
	advapi32dll = syscall.NewLazyDLL("advapi32.dll")

	_OpenProcessToken    = advapi32dll.NewProc("OpenProcessToken")
	_GetTokenInformation = advapi32dll.NewProc("GetTokenInformation")
)

const (
	THREAD_BASE_PRIORITY_LOWRT = 15  // value that gets a thread to LowRealtime-1
	THREAD_BASE_PRIORITY_MAX   = 2   // maximum thread base priority boost
	THREAD_BASE_PRIORITY_MIN   = -2  // minimum thread base priority boost
	THREAD_BASE_PRIORITY_IDLE  = -15 // value that gets a thread to idle

	SYNCHRONIZE              = 0x00100000
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	EVENT_MODIFY_STATE       = 0x0002
	EVENT_ALL_ACCESS         = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3)
	MUTANT_QUERY_STATE       = 0x0001
	MUTANT_ALL_ACCESS        = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | MUTANT_QUERY_STATE)

	SEMAPHORE_MODIFY_STATE = 0x0002
	SEMAPHORE_ALL_ACCESS   = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3)

	PROCESS_TERMINATE         = 0x0001
	PROCESS_CREATE_THREAD     = 0x0002
	PROCESS_VM_OPERATION      = 0x0008
	PROCESS_VM_READ           = 0x0010
	PROCESS_VM_WRITE          = 0x0020
	PROCESS_DUP_HANDLE        = 0x0040
	PROCESS_CREATE_PROCESS    = 0x0080
	PROCESS_SET_QUOTA         = 0x0100
	PROCESS_SET_INFORMATION   = 0x0200
	PROCESS_QUERY_INFORMATION = 0x0400
	// if NTDDI_VERSION >= NTDDI_VISTA
	PROCESS_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF)
	// else
	//PROCESS_ALL_ACCESS        = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFF);
	// endif

	PROCESSOR_INTEL_386     = 386
	PROCESSOR_INTEL_486     = 486
	PROCESSOR_INTEL_PENTIUM = 586
	PROCESSOR_INTEL_IA64    = 2200
	PROCESSOR_AMD_X8664     = 8664
	PROCESSOR_MIPS_R4000    = 4000 // incl R4101 & R3910 for Windows CE
	PROCESSOR_ALPHA_21064   = 21064
	PROCESSOR_PPC_601       = 601
	PROCESSOR_PPC_603       = 603
	PROCESSOR_PPC_604       = 604
	PROCESSOR_PPC_620       = 620
	PROCESSOR_HITACHI_SH3   = 10003  // Windows CE
	PROCESSOR_HITACHI_SH3E  = 10004  // Windows CE
	PROCESSOR_HITACHI_SH4   = 10005  // Windows CE
	PROCESSOR_MOTOROLA_821  = 821    // Windows CE
	PROCESSOR_SHx_SH3       = 103    // Windows CE
	PROCESSOR_SHx_SH4       = 104    // Windows CE
	PROCESSOR_STRONGARM     = 2577   // Windows CE - 0xA11
	PROCESSOR_ARM720        = 1824   // Windows CE - 0x720
	PROCESSOR_ARM820        = 2080   // Windows CE - 0x820
	PROCESSOR_ARM920        = 2336   // Windows CE - 0x920
	PROCESSOR_ARM_7TDMI     = 70001  // Windows CE
	PROCESSOR_OPTIL         = 0x494F // MSIL

	PROCESSOR_ARCHITECTURE_INTEL         = 0
	PROCESSOR_ARCHITECTURE_MIPS          = 1
	PROCESSOR_ARCHITECTURE_ALPHA         = 2
	PROCESSOR_ARCHITECTURE_PPC           = 3
	PROCESSOR_ARCHITECTURE_SHX           = 4
	PROCESSOR_ARCHITECTURE_ARM           = 5
	PROCESSOR_ARCHITECTURE_IA64          = 6
	PROCESSOR_ARCHITECTURE_ALPHA64       = 7
	PROCESSOR_ARCHITECTURE_MSIL          = 8
	PROCESSOR_ARCHITECTURE_AMD64         = 9
	PROCESSOR_ARCHITECTURE_IA32_ON_WIN64 = 10

	PROCESSOR_ARCHITECTURE_UNKNOWN = 0xFFFF
)

const (
	TOKEN_ASSIGN_PRIMARY    = 0x0001
	TOKEN_DUPLICATE         = 0x0002
	TOKEN_IMPERSONATE       = 0x0004
	TOKEN_QUERY             = 0x0008
	TOKEN_QUERY_SOURCE      = 0x0010
	TOKEN_ADJUST_PRIVILEGES = 0x0020
	TOKEN_ADJUST_GROUPS     = 0x0040
	TOKEN_ADJUST_DEFAULT    = 0x0080
	TOKEN_ADJUST_SESSIONID  = 0x0100

	TOKEN_ALL_ACCESS_P = (STANDARD_RIGHTS_REQUIRED | TOKEN_ASSIGN_PRIMARY |
		TOKEN_DUPLICATE | TOKEN_IMPERSONATE | TOKEN_QUERY |
		TOKEN_QUERY_SOURCE | TOKEN_ADJUST_PRIVILEGES | TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT)

	// if _WIN32_WINNT > 0x0400 || !defined(_WIN32_WINNT)
	TOKEN_ALL_ACCESS = TOKEN_ALL_ACCESS_P | TOKEN_ADJUST_SESSIONID

	TOKEN_READ  = (STANDARD_RIGHTS_READ | TOKEN_QUERY)
	TOKEN_WRITE = (STANDARD_RIGHTS_WRITE | TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS | TOKEN_ADJUST_DEFAULT)
	TOKEN_EXECUTE = STANDARD_RIGHTS_EXECUTE
)

type TTokenType uint32

const (
	TokenTPad TTokenType = iota + 0
	TokenPrimary
	TokenImpersonation
)

type TTokenInformationClass uint32

const (
	TokenUser TTokenInformationClass = iota + 1
	TokenGroups
	TokenPrivileges
	TokenOwner
	TokenPrimaryGroup
	TokenDefaultDacl
	TokenSource
	TokenType
	TokenImpersonationLevel
	TokenStatistics
	TokenRestrictedSids
	TokenSessionId
	TokenGroupsAndPrivileges
	TokenSessionReference
	TokenSandBoxInert
	TokenAuditPolicy
	TokenOrigin
	TokenElevationType
	TokenLinkedToken
	TokenElevation
	TokenHasRestrictions
	TokenAccessInformation
	TokenVirtualizationAllowed
	TokenVirtualizationEnabled
	TokenIntegrityLevel
	TokenUIAccess
	TokenMandatoryPolicy
	TokenLogonSid
	VMaxTokenInfoClass
)

// OpenProcessToken
func OpenProcessToken(ProcessHandle uintptr, DesiredAccess uint32, TokenHandle *uintptr) bool {
	r, _, _ := _OpenProcessToken.Call(ProcessHandle, uintptr(DesiredAccess), uintptr(unsafe.Pointer(TokenHandle)))
	return r != 0
}

// GetTokenInformation
func GetTokenInformation(TokenHandle uintptr, TokenInformationClass TTokenInformationClass, TokenInformation uintptr, TokenInformationLength uint32,
	ReturnLength *uint32) bool {
	r, _, _ := _GetTokenInformation.Call(TokenHandle, uintptr(TokenInformationClass), TokenInformation, uintptr(TokenInformationLength), uintptr(unsafe.Pointer(ReturnLength)))
	return r != 0
}

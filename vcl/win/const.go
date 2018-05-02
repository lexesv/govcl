package win

//  MessageBox() Flags
const (
	MB_OK               = 0x00000000
	MB_OKCANCEL         = 0x00000001
	MB_ABORTRETRYIGNORE = 0x00000002
	MB_YESNOCANCEL      = 0x00000003
	MB_YESNO            = 0x00000004
	MB_RETRYCANCEL      = 0x00000005

	MB_ICONHAND        = 0x00000010
	MB_ICONQUESTION    = 0x00000020
	MB_ICONEXCLAMATION = 0x00000030
	MB_ICONASTERISK    = 0x00000040
	MB_USERICON        = 0x00000080
	MB_ICONWARNING     = MB_ICONEXCLAMATION
	MB_ICONERROR       = MB_ICONHAND
	MB_ICONINFORMATION = MB_ICONASTERISK
	MB_ICONSTOP        = MB_ICONHAND
	MB_DEFBUTTON1      = 0x00000000
	MB_DEFBUTTON2      = 0x00000100
	MB_DEFBUTTON3      = 0x00000200
	MB_DEFBUTTON4      = 0x00000300
	MB_APPLMODAL       = 0x00000000
	MB_SYSTEMMODAL     = 0x00001000
	MB_TASKMODAL       = 0x00002000
	MB_HELP            = 0x00004000 // Help Button

	MB_NOFOCUS              = 0x00008000
	MB_SETFOREGROUND        = 0x00010000
	MB_DEFAULT_DESKTOP_ONLY = 0x00020000

	MB_TOPMOST    = 0x00040000
	MB_RIGHT      = 0x00080000
	MB_RTLREADING = 0x00100000

	MB_SERVICE_NOTIFICATION      = 0x00200000
	MB_SERVICE_NOTIFICATION_NT3X = 0x00040000

	MB_TYPEMASK = 0x0000000F
	MB_ICONMASK = 0x000000F0
	MB_DEFMASK  = 0x00000F00
	MB_MODEMASK = 0x00003000
	MB_MISCMASK = 0x0000C000
)

const (
	// Registry.  Reserved Key Handles.

	HKEY_CLASSES_ROOT     = 0x80000000
	HKEY_CURRENT_USER     = 0x80000001
	HKEY_LOCAL_MACHINE    = 0x80000002
	HKEY_USERS            = 0x80000003
	HKEY_PERFORMANCE_DATA = 0x80000004
	HKEY_CURRENT_CONFIG   = 0x80000005
	HKEY_DYN_DATA         = 0x80000006
)

const (
	SYNCHRONIZE = 0x00100000
	// The following are masks for the predefined standard access types

	DELETE                  = 0x00010000 // Renamed from DELETE
	READ_CONTROL            = 0x00020000
	WRITE_DAC               = 0x00040000
	WRITE_OWNER             = 0x00080000
	STANDARD_RIGHTS_READ    = READ_CONTROL
	STANDARD_RIGHTS_WRITE   = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE = READ_CONTROL
	STANDARD_RIGHTS_ALL     = 0x001F0000
	SPECIFIC_RIGHTS_ALL     = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY  = 0x01000000
	MAXIMUM_ALLOWED         = 0x02000000
	GENERIC_READ            = 0x80000000
	GENERIC_WRITE           = 0x40000000
	GENERIC_EXECUTE         = 0x20000000
	GENERIC_ALL             = 0x10000000

	// Registry Specific Access Rights.

	KEY_QUERY_VALUE        = 0x0001
	KEY_SET_VALUE          = 0x0002
	KEY_CREATE_SUB_KEY     = 0x0004
	KEY_ENUMERATE_SUB_KEYS = 0x0008
	KEY_NOTIFY             = 0x0010
	KEY_CREATE_LINK        = 0x0020
	KEY_WOW64_32KEY        = 0x0200
	KEY_WOW64_64KEY        = 0x0100
	KEY_WOW64_RES          = 0x0300

	// (STANDARD_RIGHTS_READ | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) & ^SYNCHRONIZE
	KEY_READ = 0x00020019
	// (STANDARD_RIGHTS_WRITE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) & ^SYNCHRONIZE
	KEY_WRITE = 0x00020006
	//  KEY_READ & ^SYNCHRONIZE
	KEY_EXECUTE = 0x00020019

	// (STANDARD_RIGHTS_ALL | KEY_QUERY_VALUE |
	//		KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS |
	//		KEY_NOTIFY | KEY_CREATE_LINK) & ^SYNCHRONIZE
	KEY_ALL_ACCESS = 0x000F003F
)

const (
	PROCESSOR_ARCHITECTURE_AMD64 = 9
	VER_NT_WORKSTATION           = 0x0000001
	SM_SERVERR2                  = 89
)

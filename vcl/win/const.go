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
	SM_SERVERR2 = 89
)

// SetWindowLongPtr GetWindowLongPtr
const (
	GWL_WNDPROC    = -4
	GWL_HINSTANCE  = -6
	GWL_HWNDPARENT = -8
	GWL_STYLE      = -16
	GWL_EXSTYLE    = -20
	GWL_USERDATA   = -21
	GWL_ID         = -12
)

// Windows Messages
const (
	WM_SYSCOMMAND = 0x0112
)

const (
	// System Menu Command Values
	SC_SIZE         = 61440
	SC_MOVE         = 61456
	SC_MINIMIZE     = 61472
	SC_MAXIMIZE     = 61488
	SC_NEXTWINDOW   = 61504
	SC_PREVWINDOW   = 61520
	SC_CLOSE        = 61536
	SC_VSCROLL      = 61552
	SC_HSCROLL      = 61568
	SC_MOUSEMENU    = 61584
	SC_KEYMENU      = 61696
	SC_ARRANGE      = 61712
	SC_RESTORE      = 61728
	SC_TASKLIST     = 61744
	SC_SCREENSAVE   = 61760
	SC_HOTKEY       = 61776
	SC_DEFAULT      = 61792
	SC_MONITORPOWER = 61808
	SC_CONTEXTHELP  = 61824
	SC_SEPARATOR    = 61455

	SCF_ISSECURE = 0x00000001

	// Obsolete names
	SC_ICON = SC_MINIMIZE
	SC_ZOOM = SC_MAXIMIZE
)

const MAX_PATH = 260

const (
	// Scroll Bar Constants
	SB_HORZ = 0
	SB_VERT = 1
	SB_CTL  = 2
	SB_BOTH = 3

	// Scroll Bar Commands
	SB_LINEUP        = 0
	SB_LINELEFT      = 0
	SB_LINEDOWN      = 1
	SB_LINERIGHT     = 1
	SB_PAGEUP        = 2
	SB_PAGELEFT      = 2
	SB_PAGEDOWN      = 3
	SB_PAGERIGHT     = 3
	SB_THUMBPOSITION = 4
	SB_THUMBTRACK    = 5
	SB_TOP           = 6
	SB_LEFT          = 6
	SB_BOTTOM        = 7
	SB_RIGHT         = 7
	SB_ENDSCROLL     = 8

	// ShowWindow() Commands
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	SW_NORMAL          = 1
	SW_SHOWMINIMIZED   = 2
	SW_SHOWMAXIMIZED   = 3
	SW_MAXIMIZE        = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
	SW_MAX             = 11

	// Old ShowWindow() Commands
	HIDE_WINDOW         = 0
	SHOW_OPENWINDOW     = 1
	SHOW_ICONWINDOW     = 2
	SHOW_FULLSCREEN     = 3
	SHOW_OPENNOACTIVATE = 4

	// Identifiers for the WM_SHOWWINDOW message
	SW_PARENTCLOSING = 1
	SW_OTHERZOOM     = 2
	SW_PARENTOPENING = 3
	SW_OTHERUNZOOM   = 4
)

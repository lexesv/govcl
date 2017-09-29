package api

type TPoint struct {
	X, Y int32
}

type TRect struct {
	Left, Top, Width, Height int32
}

type TSize struct {
	Cx, Cy int32
}

type HWND uintptr

type HBITMAP uintptr

type TModalResult int32

type HMENU uintptr

type HICON uintptr

type HDC uintptr

type TColor uint32

type TTabOrder int16

type TProgressRange int32

type TBorderWidth int32

type HFONT uintptr

type TFontStyles int32

type TImageIndex int32

type HBRUSH uintptr

type HPEN uintptr

type TFontCharset uint8

type TDateTime float64

type HKEY uintptr

type HMONITOR uintptr

type TAnchors uint8

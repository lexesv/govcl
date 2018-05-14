package rtl

import (
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/types"
)

// Move Delphi中的内存操作，不过这里传入的是指针
func Move(src, dest uintptr, llen int) {
	api.DMove(src, dest, llen)
}

// Include Delphi集合加法，val...中存储为位的索引，下标为0
func Include(r uint32, val ...uint8) uint32 {
	for _, v := range val {
		r |= (1 << uint8(v))
	}
	return r
}

// Exclude Delphi集合减法，val...中存储为位的索引，下标为0
func Exclude(r uint32, val ...uint8) uint32 {
	for _, v := range val {
		r &= ^(1 << uint8(v))
	}
	return r
}

// InSets Delphi集合类型的判断,类型，然后后面是第几位，下标为0
func InSets(r, s uint32) bool {
	if r&(1<<uint8(s)) != 0 {
		return true
	}
	return false
}

// GetStringArrOf 从一个Delphi字符串数组获取成员
func GetStringArrOf(p uintptr, index int) string {
	return api.DGetStringArrOf(p, index)
}

// SetFormCanClose 窗口的TCloseQuery事件
//func SetFormCanClose(p uintptr, val bool) {
//	a := (*int32)(unsafe.Pointer(p))
//	*a = 1
//	if !val {
//		*a = 0
//	}
//}

//// GetKey 获取var key：Char的 Unicode char size=2
//func GetKey(p uintptr) uint16 {
//	return *(*uint16)(unsafe.Pointer(p))
//}

//// SetKey 获取var key：Char的 Unicode char size=2
//func SetKey(p uintptr, val uint16) {
//	*(*uint16)(unsafe.Pointer(p)) = val
//}

// TextToShortCut 将字符串转为ShortCut类型
func TextToShortCut(val string) types.TShortCut {
	return api.DTextToShortCut(val)
}

// ShortCutToText 将ShortCut类型转为字符串
func ShortCutToText(val types.TShortCut) string {
	return api.DShortCutToText(val)
}

// SysOpen 打开，windows下调用ShellExecute
func SysOpen(filename string) {
	api.DSysOpen(filename)
}

// ExtractFilePath 提取文件名的路径，带“\”的
func ExtractFilePath(filename string) string {
	return api.DExtractFilePath(filename)
}

// FileExists 判断文件是否存在
func FileExists(filename string) bool {
	return api.DFileExists(filename)
}

// InheritsFromControl 判断对象是否继承自TControl
func InheritsFromControl(obj uintptr) bool {
	return api.DInheritsFromControl(obj)
}

// InheritsFromWinControl 判断对象是否继承自TWinControl
func InheritsFromWinControl(obj uintptr) bool {
	return api.DInheritsFromWinControl(obj)
}

// InheritsFromComponent 判断对象是否继承自TComponent
func InheritsFromComponent(obj uintptr) bool {
	return api.DInheritsFromComponent(obj)
}

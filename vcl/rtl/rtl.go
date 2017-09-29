package rtl

import (
	"unsafe"

	"gitee.com/ying32/govcl/vcl/api"
	. "gitee.com/ying32/govcl/vcl/types"
)

// Move Delphi中的内存操作，不过这里传入的是指针
func Move(src, dest uintptr, llen int) {
	api.DMove(src, dest, llen)
}

// MainInstance EXE自身的实例
func MainInstance() uintptr {
	return api.DGetMainInstance()
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

// SetFormCanClose 窗口的TCloseQuery事件
func SetFormCanClose(p uintptr, val bool) {
	a := (*int32)(unsafe.Pointer(p))
	*a = 1
	if !val {
		*a = 0
	}
}

// GetKey 获取var key：Char的 Unicode char size=2
func GetKey(p uintptr) uint16 {
	return *(*uint16)(unsafe.Pointer(p))
}

// SetKey 获取var key：Char的 Unicode char size=2
func SetKey(p uintptr, val uint16) {
	*(*uint16)(unsafe.Pointer(p)) = val
}

// SetReportMemoryLeaksOnShutdown 程序结束时报告内存泄露，总有2-4字节的未知泄露位置
func SetReportMemoryLeaksOnShutdown(v bool) {
	api.DSetReportMemoryLeaksOnShutdown(v)
}

// TextToShortCut 将字符串转为ShortCut类型
func TextToShortCut(val string) TShortCut {
	return api.DTextToShortCut(val)
}

// ShortCutToText 将ShortCut类型转为字符串
func ShortCutToText(val TShortCut) string {
	return api.DShortCutToText(val)
}

// SysOpen 打开，windows下调用ShellExecute
func SysOpen(filename string) {
	api.DSysOpen(filename)
}

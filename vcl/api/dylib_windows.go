// 实现一个动态调用call的，主要是为了解决异常问题
package api

import (
	"fmt"
	"strconv"
	"syscall"
)

type LazyDLL struct {
	*syscall.LazyDLL
	// 加了异常处理的
	mySyscall *syscall.LazyProc
}

// 因为lcl中在windows下与原vcl中使用的字符编码不一致，当加载了lcl库后会有乱码的
var IsloadedLcl = false

func NewLazyDLL(name string) *LazyDLL {
	l := new(LazyDLL)
	l.LazyDLL = syscall.NewLazyDLL(name)
	// 这里做个判断，当libvcl.dll或者libvclx64.dll加载失败时尝试加载liblcl.dll
	// 这样做主要为以后考虑，对于某些人来说怕什么的来说，可以使用非Delphi的组件
	if l.Load() != nil {
		fmt.Println(fmt.Sprintf("\"%s\" does not exist, trying to load liblcl.dll.", name))
		l.LazyDLL = syscall.NewLazyDLL("liblcl.dll")
		if l.Load() == nil {
			IsloadedLcl = true
		}
	}
	// 导入调用的
	l.mySyscall = l.LazyDLL.NewProc("MySyscall")
	if l.mySyscall.Find() != nil {
		l.mySyscall = nil
	}
	return l
}

// 定义一个相同的
type LazyProc struct {
	lzProc *syscall.LazyProc
	lzdll  *LazyDLL
}

func (d *LazyDLL) NewProc(name string) *LazyProc {
	l := new(LazyProc)
	l.lzProc = d.LazyDLL.NewProc(name)
	l.lzdll = d
	return l
}

func (d *LazyDLL) call(proc *LazyProc, a ...uintptr) (r1, r2 uintptr, lastErr error) {
	// 没到找到我封装的那个系统函数，就使用原始的
	if d.mySyscall == nil {
		return proc.CallOriginal(a...)
	}
	err := proc.Find()
	if err != nil {
		fmt.Println("proc \"" + proc.lzProc.Name + "\" not find.")
		return 0, 0, syscall.EINVAL
	}
	addr := proc.Addr()
	if addr != 0 {
		pLen := uintptr(len(a))
		switch pLen {
		case 0:
			return d.mySyscall.Call(addr, pLen)
		case 1:
			return d.mySyscall.Call(addr, pLen, a[0])
		case 2:
			return d.mySyscall.Call(addr, pLen, a[0], a[1])
		case 3:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2])
		case 4:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3])
		case 5:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4])
		case 6:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5])
		case 7:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6])
		case 8:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
		case 9:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
		case 10:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9])
		case 11:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10])
		case 12:
			return d.mySyscall.Call(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
		default:
			panic("Call " + proc.lzProc.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
		}
	}
	return 0, 0, syscall.EINVAL
}

func (p *LazyProc) Addr() uintptr {
	return p.lzProc.Addr()
}

func (p *LazyProc) Find() error {
	return p.lzProc.Find()
}

func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	return p.lzdll.call(p, a...)
}

func (p *LazyProc) CallOriginal(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	return p.lzProc.Call(a...)
}

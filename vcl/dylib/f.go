// +build linux,cgo darwin,cgo

package dylib

/*
	#cgo LDFLAGS: -ldl

    #cgo LDFLAGS: -L. -Wl,-rpath,${SRCDIR} -llibvcl
	#cgo LDFLAGS: -L. -Wl,-rpath,${SRCDIR} -lcgunwind.1.0
	#cgo LDFLAGS: -L. -Wl,-rpath,${SRCDIR} -lcrossvcl32
	#cgo LDFLAGS: -L. -Wl,-rpath,${SRCDIR} -lcrossvcl


	// crossvcl32
	// crossvcl
	// libvcl

	#include <dlfcn.h>
	#include <limits.h>
	#include <stdlib.h>
	#include <stdint.h>

	#include <stdio.h>

	static uintptr_t pluginOpen(const char* path) {
	     void* h = dlopen(path, RTLD_NOW|RTLD_GLOBAL); // RTLD_LAZY RTLD_NOW |RTLD_GLOBAL
	     if (h == NULL) {
		     printf("dlopen err: %s\n", (char*)dlerror());
	         return 0;
	     }
	     return (uintptr_t)h;
	}

	static uintptr_t pluginLookup(uintptr_t h, const char* name) {
	     void* r = dlsym((void*)h, name);
	     if (r == NULL) {
			 printf("dlsym err: %s\n", (char*)dlerror());
	         return 0;
	     }
	     return (uintptr_t)r;
	}

	static void pluginClose(uintptr_t h) {
	     if(h != 0) {
	         dlclose((void*)h);
	     }
	}

    extern void* Syscall0(void*);
	extern void* Syscall1(void*,void*);
	extern void* Syscall2(void*,void*,void*);
	extern void* Syscall3(void*,void*,void*,void*);
	extern void* Syscall4(void*,void*,void*,void*,void*);
	extern void* Syscall5(void*,void*,void*,void*,void*,void*);
	extern void* Syscall6(void*,void*,void*,void*,void*,void*,void*);
*/
import "C"

import (
	"fmt"

	"strconv"
	"sync"
	"syscall"
	"unsafe"
)

type LazyDLL struct {
	mu     sync.Mutex
	handle C.uintptr_t
	Name   string
}

func NewLazyDLL(name string) *LazyDLL {

	m := new(LazyDLL)
	m.Name = name
	m.mu.Lock()
	defer m.mu.Unlock()

	cRelName := C.CString(name)
	defer C.free(unsafe.Pointer(cRelName))
	m.handle = C.pluginOpen(cRelName)
	if m.handle == 0 {
		panic("dlopen(" + name + ")， 失败。")
	}
	return m
}

func (l *LazyDLL) NewProc(name string) *LazyProc {
	p := new(LazyProc)
	p.Name = name
	cRelName := C.CString(name)
	defer C.free(unsafe.Pointer(cRelName))
	p.p = uintptr(C.pluginLookup(l.handle, cRelName))
	return p
}

func (l *LazyDLL) Close() {
	C.pluginClose(l.handle)
}

type LazyProc struct {
	p    uintptr
	Name string
}

func (p *LazyProc) Addr() uintptr {
	return p.p
}

func toPtr(a uintptr) unsafe.Pointer {
	return unsafe.Pointer(a)
}

func (p *LazyProc) Call(a ...uintptr) (uintptr, uintptr, syscall.Errno) {
	fmt.Println("name:", p.Name, ", p.p=", p.p, ", args: ", a)

	if p.p == 0 {
		return 0, 0, syscall.EV_ERROR
	}
	var ret unsafe.Pointer
	switch len(a) {
	case 0:
		ret = C.Syscall0(unsafe.Pointer(p.p))
	case 1:
		ret = C.Syscall1(toPtr(p.p), toPtr(a[0]))
	case 2:
		ret = C.Syscall2(toPtr(p.p), toPtr(a[0]), toPtr(a[1]))
	case 3:
		ret = C.Syscall3(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]))
	case 4:
		ret = C.Syscall4(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]))
	case 5:
		ret = C.Syscall5(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]))
	case 6:
		ret = C.Syscall6(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]))
	default:
		panic("Call " + p.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
	}
	return uintptr(ret), 0, 0

}

//go:build gc && !go1.19

package ohmem

import (
	"reflect"
	"unsafe"
)

// sysAlloc allocates memory off heap by calling runtime.sysAlloc.
//
// DON'T use this if you don't know what it does.
func sysAlloc(n uintptr) []byte {
	var memStat uint64
	addr := runtime_sysAlloc(n, &memStat)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(addr),
		Len:  int(n),
		Cap:  int(n),
	}))
}

// sysFree frees memory allocated by sysAlloc.
//
// DON'T use this if you don't know what it does.
func sysFree(mem []byte) {
	addr := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&mem)).Data)
	n := uintptr(cap(mem))
	memStat := uint64(n)
	runtime_sysFree(addr, n, &memStat)
}

//go:linkname runtime_sysAlloc runtime.sysAlloc
func runtime_sysAlloc(n uintptr, sysStat *uint64) unsafe.Pointer

//go:linkname runtime_sysFree runtime.sysFree
func runtime_sysFree(v unsafe.Pointer, n uintptr, sysStat *uint64)

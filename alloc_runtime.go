//go:build !mimalloc

package ohmem

func _C_zalloc(n int) []byte {
	mem := sysAlloc(uintptr(n))
	for i := range mem {
		mem[i] = 0
	}
	return mem
}

func _C_free(mem []byte) {
	sysFree(mem)
}

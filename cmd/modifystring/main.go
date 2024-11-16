package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	str := strings.Clone("аб")
	bytes := []byte(str)
	fmt.Println(bytes)

	ptr := unsafe.Pointer(unsafe.StringData(str))

	cur := (*byte)(unsafe.Pointer(uintptr(ptr) + 1))
	*cur = 177
	fmt.Println(str)
}

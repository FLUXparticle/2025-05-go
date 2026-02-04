package main

/*
//Falls die lib extern compiliert wird
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./build -lmylib
#include "mylib.h"
*/
import "C"

import "fmt"

func main() {
	erg := C.add(1, 2)
	fmt.Println(int(erg))
}

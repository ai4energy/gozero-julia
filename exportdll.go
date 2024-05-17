package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"github.com/ai4energy/gozero-julia/cmd"
)

//export JuliaGenCode
func JuliaGenCode(outputDir, apiFile *C.char) *C.char {
	err := cmd.JuliaGenCode(C.GoString(outputDir), C.GoString(apiFile))
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("Success")
}

func main() {}

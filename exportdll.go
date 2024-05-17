package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"github.com/ai4energy/gozero-julia/cmd"
	"github.com/ai4energy/gozero-julia/prepare"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"sync"
)

var (
	jsonData string
	mu       sync.Mutex
)

//export JuliaGenCode
func JuliaGenCode(outputDir, apiFile *C.char) *C.char {
	err := cmd.JuliaGenCode(C.GoString(outputDir), C.GoString(apiFile))
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("Success")
}

// 将 ApiSpec 转换为 JSON 并存储在全局变量中
func setJSON(apiSpec *spec.ApiSpec) error {
	data, err := json.Marshal(apiSpec)
	if err != nil {
		return err
	}
	mu.Lock()
	jsonData = string(data)
	mu.Unlock()
	return nil
}

//export ExportToJSON
func ExportToJSON(outputDir, apiFile *C.char) *C.char {
	// 设置 ApiFile
	prepare.ApiFile = C.GoString(apiFile)
	prepare.OutputDir = C.GoString(outputDir)
	prepare.Setup()

	err := setJSON(prepare.ApiSpec)
	if err != nil {
		return C.CString(err.Error())
	}

	// 返回 JSON 数据
	mu.Lock()
	defer mu.Unlock()
	return C.CString(jsonData)
}

func main() {}

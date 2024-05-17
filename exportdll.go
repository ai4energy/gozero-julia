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
)

//export JuliaGenCode
func JuliaGenCode(outputDir, apiFile *C.char) *C.char {
	err := cmd.JuliaGenCode(C.GoString(outputDir), C.GoString(apiFile))
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("Success")
}

// 将 ApiSpec 转换为 JSON 并返回
func convertToJSON(apiSpec *spec.ApiSpec) (string, error) {
	data, err := json.Marshal(apiSpec)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//export ExportToJSON
func ExportToJSON(apiFile *C.char) *C.char {
	// 设置 ApiFile
	prepare.ApiFile = C.GoString(apiFile)
	prepare.Setup()

	jsonData, err := convertToJSON(prepare.ApiSpec)
	if err != nil {
		return C.CString(err.Error())
	}

	// 返回 JSON 数据
	return C.CString(jsonData)
}

func main() {}

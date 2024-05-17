package generator

import (
	"fmt"

	"github.com/ai4energy/gozero-julia/prepare"
)

func GenerateAll() error {
	if err := prepare.PrepareOutputDir(prepare.OutputDir); err != nil {
		return err
	}

	rootPkg, err := prepare.GetParentPackage(prepare.OutputDir)
	if err != nil {
		return err
	}
	prepare.RootPkg = rootPkg

	// 生成代码的逻辑
	fmt.Println("Generating...") // 输出生成中的提示信息
	fmt.Println(prepare.ApiSpec.Service.Name)

	return nil
}

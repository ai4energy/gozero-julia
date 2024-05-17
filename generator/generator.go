package generator

import (
	"fmt"
	"github.com/ai4energy/gozero-julia/prepare"
)

func GenerateAll() error {
	fmt.Println("Generating...") // 输出生成中的提示信息
	fmt.Println(prepare.ApiSpec.Service.Name)

	return nil

}

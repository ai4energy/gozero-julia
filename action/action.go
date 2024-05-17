package action

import (
	"fmt"
	"github.com/ai4energy/gozero-julia/generator"
	"github.com/ai4energy/gozero-julia/prepare"
	"github.com/urfave/cli/v2"
)

func JuliaGenCode(ctx *cli.Context) error {
	// 从命令行参数中获取值
	prepare.OutputDir = ctx.String("dir")
	prepare.ApiFile = ctx.String("api")

	// 打印 Hello, world 和目录信息
	fmt.Println("Hello, world!")
	fmt.Println("Output Directory:", prepare.OutputDir)
	fmt.Println("API File:", prepare.ApiFile)
	prepare.Setup()
	Must(generator.GenerateAll())
	return nil
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

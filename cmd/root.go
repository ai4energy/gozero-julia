package cmd

import (
	"fmt"
	"github.com/ai4energy/gozero-julia/generator"
	"github.com/ai4energy/gozero-julia/prepare"
	"github.com/spf13/cobra"
	"os"
)

var (
	outputDir string
	apiFile   string
	rootCmd   = &cobra.Command{
		Use:     "gozero-julia",
		Short:   "生成基于 Julia(Oxygen) 框架的 WEB 服务的极简项目结构",
		Example: "gozero-julia --dir=. some.api",
		Args:    cobra.NoArgs,
		RunE:    RunJuliaGenCode,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil ./{
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&outputDir, "dir", ".", "生成项目目录")
	rootCmd.Flags().StringVar(&apiFile, "api", "", "API文件路径")
}

func RunJuliaGenCode(cmd *cobra.Command, args []string) error {
	return JuliaGenCode(outputDir, apiFile)
}

func JuliaGenCode(outputDir string, apiFile string) error {
	prepare.OutputDir = outputDir
	prepare.ApiFile = apiFile

	// Print Hello, world and directory information
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

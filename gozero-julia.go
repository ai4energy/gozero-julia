package main

import (
	"fmt"
	"github.com/ai4energy/gozero-julia/action"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

var (
	version  = "20240517"
	commands = []*cli.Command{
		{
			Name:   "julia",
			Usage:  "generates http client for julia",
			Action: action.JuliaGenCode,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "package",
					Usage: "the package of julia",
				},
				&cli.StringFlag{
					Name:    "dir",
					Usage:   "directory to generate the project",
					Value:   ".", // 默认值为当前目录
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:  "api",
					Usage: "api file to generate the julia client from",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "gozero-julia"
	app.Usage = "a plugin of goctl to generate http client code for julia"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error running gozero-julia: %+v\n", err)
	}
}

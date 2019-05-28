package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hapiman/remember/lib"
	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(child),
		cli.Tree(ls),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help information")

// root command
type rootT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

var root = &cli.Command{
	Desc: "this is a reminder command of notes",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}

// child command
type childT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
	File string `cli:"file" usage:"will view file name"`
	Dir  string `cli:"dir" usage:"view which directory"`
}

var child = &cli.Command{
	Name: "show",
	Desc: "view file content",
	Argv: func() interface{} { return new(childT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childT)
		fl := argv.File
		if fl == "" {
			log.Fatal("缺少文件名称")
		}
		docDir := computeWorkDir(argv)
		oneS := fmt.Sprintf("%s/%s", docDir, fl)
		//lib判断文件是否存在
		flag, _ := lib.PathExists(oneS)
		if flag == false {
			log.Fatal("文件不存在")
		}
		content := lib.OutputContent(oneS)
		ctx.String("the file content: \n %s \n", content)
		return nil
	},
}

var ls = &cli.Command{
	Name: "ls",
	Desc: "list all filenames",
	Argv: func() interface{} { return new(childT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childT)
		docDir := computeWorkDir(argv)
		filenames := lib.ViewFiles(docDir)
		for idx, element := range filenames {
			ctx.String("idx: %d, filenames: %s \n", idx, element)
		}
		return nil
	},
}

func computeWorkDir(argv *childT) string {
	if argv.Dir != "" {
		return argv.Dir
	}
	curDir, _ := os.Getwd()
	return fmt.Sprintf("%s/docs", curDir)
}

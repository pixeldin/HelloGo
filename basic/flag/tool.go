package main

import (
	"HelloGo/basic/flag/sun"
	"HelloGo/basic/flag/wind"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//内置flag实现
	//BuiltInFlag()
	//SubCommands()

	//自定义命令行工具
	SelfDefineFlag()
}

func AllRight() bool {
	return true
}

func BuiltInFlag()  {
	//声明接收变量, 注意返回的是个指针类型
	wordParam := flag.String("word", "defaultValue", "Desc param 1 for something.")
	numParam := flag.Int("number", 0, "Desc param 2 for integer number.")

	//an existing var declared elsewhere in the program
	ok := AllRight()

	flag.BoolVar(&ok, "ok", false, "Desc value for boolean.")

	flag.Parse()

	log.Printf("Flag for wordParam: %s", *wordParam)
	log.Printf("Flag for wordParam: %d", *numParam)
	log.Printf("Flag for inner bool value : %v", ok)
}

func SubCommands() {
	//param1: 命令参数, param2: 错误退出码
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	//声明子命令fc, fc2
	fc := fooCmd.String("fc", "default of fc", "foo sub value1")
	fc2 := fooCmd.String("fc2", "default of fc2", "foo sub value2")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	//声明子命令bc
	bc := barCmd.String("bc", "default of bc", "bar sub b")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		fooCmd.Usage()
		barCmd.Usage()
		os.Exit(1)
	}

	//参数逻辑调用, 使用分支语句
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		log.Printf("fc of foo: %s", *fc)
		log.Printf("fc2 of foo: %s", *fc2)
	case "bar":
		barCmd.Parse(os.Args[2:])
		log.Printf("bc of bar: %s", *bc)
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}

/*
	自定义方法调用
 */
//封装命令结构, 指令/执行方法/提示
type cmd struct {
	Name    string
	Process func(args ...string)
	Usage   func() string
}

//全局指令数组
var all []*cmd

func init()  {
	all = append(all, &cmd{"sun", sun.Process, sun.Usage})
	all = append(all, &cmd{"wind", wind.Process, wind.Usage})
}

//全局用法
func usage() string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("Usage: %s <command> [args...]\n", os.Args[0]))
	for _, c := range all {
		sb.WriteString("\n命令: ")
		sb.WriteString(c.Name)
		sb.WriteString(", 用法: ")
		sb.WriteString(c.Usage())
		sb.WriteString("\n")
	}
	return sb.String()
}

func SelfDefineFlag()  {
	if len(os.Args) > 1 && os.Args[1] != "--help" && os.Args[1] != "-h"{
		for _, c := range all {
			//匹配全局命令
			if c.Name == os.Args[1] {
				c.Process(os.Args[2:]...)
				return
			}
		}
		fmt.Fprintf(os.Stderr, "No match func for %s.\n", os.Args[1])
	}
	fmt.Fprintln(os.Stderr, usage())
}
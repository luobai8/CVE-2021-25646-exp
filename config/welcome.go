package config

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func Welcome() string {

	// 判断是否输入了参数
	if len(os.Args) < 2 || os.Args[1] != "-t" {
		fmt.Println("")
		fmt.Println(color.RedString("			<--- 参数输入错误 --->"))
		os.Exit(0)
	} else if len(os.Args) < 3 {

		// fmt.Println("		                可同时输入多个检测目标，需要以逗号分隔")
		// fmt.Println("			eg：infoscan.exe -i baidu.com,192.168.169.10,qq.com")
		fmt.Println("")
		fmt.Println(color.RedString("			<--- 未给参数传值，请输入目标 --->"))
		os.Exit(0)
	}

	// 接收第二个参数,作为目标的值
	target := os.Args[2]
	if strings.Contains(target, "/") || strings.Contains(target, "http") || strings.Contains(target, "https") {
		fmt.Println(color.RedString("			<--- 目标格式输入输入错误 --->"))
		os.Exit(0)
	}

	return target

}

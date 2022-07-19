/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 17:06:36
 * @LastEditTime: 2022-07-19 17:34:40
 * @LastEditors: NyanCatda
 * @Description: 打印Help信息
 * @FilePath: \NitoriConsole\Help\Print.go
 */
package Help

import (
	"fmt"

	"github.com/nyancatda/NitoriConsole/Command"
)

// 自定义打印函数
var PrintFunc func(string)

/**
 * @description: 打印Help信息
 * @param {func(string)} PrintFunc 打印函数
 * @return {*}
 */
func Print() {
	// 如果没有自定义打印函数则使用默认打印函数
	if PrintFunc == nil {
		PrintFunc = func(Command string) {
			fmt.Println(Command)
		}
	}

	PrintFunc("---------------- Help ----------------")

	// 打印Help信息
	for Command, CommandInfo := range Command.CommandList {
		// 跳过help命令
		if Command == "help" || Command == "?" {
			continue
		}

		PrintFunc(Command + " " + CommandInfo.Help)
	}

	PrintFunc("--------------------------------------")
}

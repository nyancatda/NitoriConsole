/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 17:06:36
 * @LastEditTime: 2022-07-19 17:20:22
 * @LastEditors: NyanCatda
 * @Description: 打印Help信息
 * @FilePath: \NitoriConsole\Help\Print.go
 */
package Help

import "github.com/nyancatda/NitoriConsole/Command"

/**
 * @description: 打印Help信息
 * @param {func(string)} PrintFunc 打印函数
 * @return {*}
 */
func Print(PrintFunc func(string)) {
	PrintFunc("---------------- Help ----------------")

	// 打印Help信息
	for Command, CommandInfo := range Command.CommandList {
		PrintFunc(Command + " " + CommandInfo.Help)
	}

	PrintFunc("--------------------------------------")
}

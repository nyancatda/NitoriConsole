/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 15:49:39
 * @LastEditTime: 2022-07-19 16:04:53
 * @LastEditors: NyanCatda
 * @Description: 控制台模块
 * @FilePath: \NitoriConsole\Console.go
 */
package NitoriConsole

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nyancatda/NitoriConsole/Command"
)

/**
 * @description: 启动控制台
 * @param {string} Prompt 提示符号，例如$
 * @param {func(error)} Error 错误回调函数
 * @return {*}
 */
func Start(Prompt string, Error func(error)) {
	Reader := bufio.NewReader(os.Stdin)
	// 循环处理输入
	for {
		// 打印提示符号
		fmt.Printf("\r" + Prompt)

		// 获取输入
		CmdString, err := Reader.ReadString('\n')
		if err != nil {
			Error(err)
		}

		// 执行命令
		err = Command.Execution(CmdString)
		if err != nil {
			Error(err)
		}
	}
}

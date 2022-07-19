package Help

import "github.com/nyancatda/NitoriConsole/Command"

/**
 * @description: 注册Help命令
 * @return {*}
 */
func Register() {
	// 注册Help命令
	Command.Add("help", "查看帮助", func(_ string) {
		Print()
	})
	Command.Add("?", "查看帮助", func(_ string) {
		Print()
	})
}
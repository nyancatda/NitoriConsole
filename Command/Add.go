/*
 * @Author: NyanCatda
 * @Date: 2022-07-19 16:05:50
 * @LastEditTime: 2022-07-19 16:55:44
 * @LastEditors: NyanCatda
 * @Description: 添加命令模块
 * @FilePath: \NitoriConsole\Command\Add.go
 */
package Command

/**
 * @description: 注册一条命令
 * @param {Command} string 命令名称
 * @param {Help} string 命令帮助信息
 * @param {Callback} func(string) 命令回调函数
 * @return {*}
 */
func Add(Command string, Help string, Callback func(string)) {
	// 添加至命令列表
	CommandList = make(map[string]CommandInfo)
	CommandList[Command] = CommandInfo{Help, Callback}
	return
}

# NitoriConsole
Golang控制台库

通过数行代码即可快速的创建一个带有控制台的程序

## 🎬 如何使用
### 安装
```
go get -u github.com/nyancatda/NitoriConsole
```
### 例子
``` go
package main

import (
	"fmt"

	"github.com/nyancatda/NitoriConsole"
	"github.com/nyancatda/NitoriConsole/Command"
)

func main() {
	// 注册命令
	Command.Add("say", "Say String", func(CommandStr string) {
		// 解析命令参数
		_, CommandBody := Command.Parse(CommandStr)

		// 打印参数
		fmt.Println(CommandBody[0])
	})

	// 启动控制台
	NitoriConsole.Start("$ ", false, func(err error) {
		fmt.Println(err)
	})
}

```
### 效果
``` shell
go run .

$ help
---------------- Help ----------------
say Say String
--------------------------------------
$ say suki
suki
```
### 配合AyaLog使用
推荐使用以下配置，配合[AyaLog](https://github.com/nyancatda/AyaLog)使用

``` go
package main

import (
	"time"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/NitoriConsole"
	"github.com/nyancatda/NitoriConsole/Command"
	"github.com/nyancatda/NitoriConsole/Help"
)

func main() {
	// 配置AyaLog
	AyaLog.CleanLine = true // 清理行前内容
	AyaLog.LineEndString = "$ " // 行末增加命令符

	// 使用AyaLog输出Help
	Help.PrintFunc = func(Text string) {
		AyaLog.Info("System", Text)
	}

	// 注册命令
	Command.Add("say", "Say String", func(CommandStr string) {
		// 解析命令参数
		_, CommandBody := Command.Parse(CommandStr)

		// 打印参数
		AyaLog.Info("Say", CommandBody[0])
	})

	// 延迟三秒打印信息
	go func() {
		time.Sleep(time.Second * 3)
		AyaLog.Info("System", "Hello World!")
	}()

	// 启动控制台
	NitoriConsole.Start("$ ", false, func(err error) {
		AyaLog.Error("System", err)
	})
}
```

## 📖 许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档
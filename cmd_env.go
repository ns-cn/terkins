package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
)

var CmdEnv = goter.Command{Cmd: &cobra.Command{
	Use:   "env",
	Short: "查看当前已配置的变量（含命令行参数）",
	Run: func(cmd *cobra.Command, args []string) {
		sysEnvHost, _ = os.LookupEnv(TERKINS_HOST)
		sysEnvUser, _ = os.LookupEnv(TERKINS_USER)
		sysEnvPass, _ = os.LookupEnv(TERKINS_PASS)
		sysEnvEncrypted, _ = os.LookupEnv(TERKINS_ENCRYPTED)
		table, err := gotable.Create("项目", "命令行参数", "系统环境变量", "默认值")
		if err != nil {
			return
		}
		_ = table.AddRow([]string{"主机", envHost, sysEnvHost, ""})
		_ = table.AddRow([]string{"用户名", envUser, sysEnvUser, ""})
		_ = table.AddRow([]string{"密码", envPass, sysEnvPass, ""})
		_ = table.AddRow([]string{"加密", envEncrypted, sysEnvEncrypted, "Y"})
		fmt.Print(table)
		fmt.Println("值的取值顺序：命令行参数 > 系统环境变量 > 默认值")
	},
}}

package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
	"terkins/env"
)

var CmdEnv = goter.Command{Cmd: &cobra.Command{
	Use:   "env",
	Short: "查看当前已配置的变量（含命令行参数）",
	Run: func(cmd *cobra.Command, args []string) {
		env.SysEnvHost, _ = os.LookupEnv(env.TERKINS_HOST)
		env.SysEnvUser, _ = os.LookupEnv(env.TERKINS_USER)
		env.SysEnvPass, _ = os.LookupEnv(env.TERKINS_PASS)
		env.SysEnvEncrypted, _ = os.LookupEnv(env.TERKINS_ENCRYPTED)
		table, err := gotable.Create("项目", "命令行参数", "系统环境变量", "默认值")
		if err != nil {
			return
		}
		_ = table.AddRow([]string{"主机", env.Host.Value, env.SysEnvHost, ""})
		_ = table.AddRow([]string{"用户名", env.User.Value, env.SysEnvUser, ""})
		_ = table.AddRow([]string{"密码", env.Pass.Value, env.SysEnvPass, ""})
		_ = table.AddRow([]string{"加密", env.Encrypt.Value, env.SysEnvEncrypted, "Y"})
		fmt.Print(table)
		fmt.Println("值的取值顺序：命令行参数 > 系统环境变量 > 默认值")
	},
}}

package main

import (
	"github.com/ns-cn/goter"
	"terkins/env"
)

func main() {
	root := goter.NewRootCmd("terkins", "tool to operate jenkins with terminal", env.VERSION)
	root.AddCommand(CmdJobs.Bind(env.Host, env.User, env.Pass, env.Encrypt, env.Debug))
	root.AddCommand(CmdJob.Bind(env.Host, env.User, env.Pass, env.Encrypt, env.Debug))
	root.AddCommand(CmdEnv.Bind(env.Host, env.User, env.Pass, env.Encrypt))
	root.AddCommand(CmdBuild.Bind(env.Host, env.User, env.Pass, env.Encrypt, env.Debug, env.ShowBuildInfo))
	root.AddCommand(CmdEncrypted.Bind(env.Pass, env.User))
	// 数据源
	_ = root.Execute()
}

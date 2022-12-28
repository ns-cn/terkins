package main

import (
	"github.com/ns-cn/goter"
)

func main() {
	root := goter.NewRootCmd("terkins", "tool to operate jenkins with terminal", VERSION)
	root.AddCommand(CmdJobs.Bind(flagHost, flagUser, flagPass, flagEncrypt, flagDebug))
	root.AddCommand(CmdJob.Bind(flagHost, flagUser, flagPass, flagEncrypt, flagDebug))
	root.AddCommand(CmdEnv.Bind(flagHost, flagUser, flagPass, flagEncrypt))
	root.AddCommand(CmdBuild.Bind(flagHost, flagUser, flagPass, flagEncrypt, flagDebug, flagInfoToBuild))
	root.AddCommand(CmdEncrypted.Bind(flagPass, flagUser))
	// 数据源
	_ = root.Execute()
}

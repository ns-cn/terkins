package main

import "github.com/ns-cn/goter"

const (
	TERKINS_HOST      = "TERKINS_HOST"
	TERKINS_USER      = "TERKINS_USER"
	TERKINS_PASS      = "TERKINS_PASS"
	TERKINS_ENCRYPTED = "TERKINS_ENCRYPTED"
)

// 通用的环境变量及参数
var (
	// effected setting
	envHost      = ""
	envUser      = ""
	envPass      = ""
	envEncrypted = ""
	envDebug     = ""

	flagHost    = goter.CmdStringFlag{P: &envHost, Name: "envHost", Shorthand: "H", Value: "", Usage: "envHost of jenkins, can use ENV TERKINS_HOST instead"}
	flagUser    = goter.CmdStringFlag{P: &envUser, Name: "envUser", Shorthand: "U", Value: "", Usage: "envUser of jenkins, can use ENV TERKINS_USER instead"}
	flagPass    = goter.CmdStringFlag{P: &envPass, Name: "password", Shorthand: "P", Value: "", Usage: "password of jenkins, can use ENV TERKINS_PASS instead"}
	flagEncrypt = goter.CmdStringFlag{P: &envEncrypted, Name: "envEncrypted", Shorthand: "E", Value: "", Usage: "password envEncrypted: Y/N , can use TERKINS_ENCRYPTED instead"}
	flagDebug   = goter.CmdStringFlag{P: &envDebug, Name: "envDebug", Shorthand: "D", Value: "N", Usage: "envDebug: Y/N"}

	isEncrypted = true
	// env from system
	sysEnvHost      = ""
	sysEnvUser      = ""
	sysEnvPass      = ""
	sysEnvEncrypted = ""
)

// 构建使用的参数
var (
	envBuildInfo    = ""
	flagInfoToBuild = goter.CmdStringFlag{P: &envBuildInfo, Name: "info", Shorthand: "I", Value: "Y", Usage: "info each job to build or not"}
)

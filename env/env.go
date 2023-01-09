package env

import "github.com/ns-cn/goter"

const (
	TERKINS_HOST      = "TERKINS_HOST"
	TERKINS_USER      = "TERKINS_USER"
	TERKINS_PASS      = "TERKINS_PASS"
	TERKINS_ENCRYPTED = "TERKINS_ENCRYPTED"
)

// 通用的环境变量及参数
var (
	Host    = goter.NewCmdFlagString("", "envHost", "H", "envHost of jenkins, can use ENV TERKINS_HOST instead")
	User    = goter.NewCmdFlagString("", "envUser", "U", "envUser of jenkins, can use ENV TERKINS_USER instead")
	Pass    = goter.NewCmdFlagString("", "password", "P", "password of jenkins, can use ENV TERKINS_PASS instead")
	Encrypt = goter.NewCmdFlagString("Y", "envEncrypted", "E", "password envEncrypted: Y/N , can use TERKINS_ENCRYPTED instead")
	Debug   = goter.NewCmdFlagBool(false, "debug", "D", "debug to show dev information")

	IsEncrypted = true
	// env from system
	SysEnvHost      = ""
	SysEnvUser      = ""
	SysEnvPass      = ""
	SysEnvEncrypted = ""
)

// 构建使用的参数
var (
	ShowBuildInfo = goter.NewCmdFlagBool(false, "info", "I", "info each job to build or not")
)

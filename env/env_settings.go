package env

import (
	"encoding/hex"
	"github.com/ns-cn/goter"
	"os"
)

func ReadSetting() {
	SysEnvHost, _ = os.LookupEnv(TERKINS_HOST)
	SysEnvUser, _ = os.LookupEnv(TERKINS_USER)
	SysEnvPass, _ = os.LookupEnv(TERKINS_PASS)
	SysEnvEncrypted, _ = os.LookupEnv(TERKINS_ENCRYPTED)
	if Host.Value == "" {
		Host.Value = SysEnvHost
	}
	if User.Value == "" {
		User.Value = SysEnvUser
	}
	if Pass.Value == "" {
		Pass.Value = SysEnvPass
	}
	if Encrypt.Value == "" {
		Encrypt.Value = SysEnvEncrypted
	}
	IsEncrypted = goter.IsYes(Encrypt.Value, true)
	if IsEncrypted {
		hexBytes, _ := hex.DecodeString(Pass.Value)
		Pass.Value = string(goter.AesDecryptCBC(hexBytes, goter.GetKey(User.Value)))
	}
}

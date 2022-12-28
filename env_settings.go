package main

import (
	"encoding/hex"
	"github.com/ns-cn/goter"
	"os"
)

func ReadSetting() {
	sysEnvHost, _ = os.LookupEnv(TERKINS_HOST)
	sysEnvUser, _ = os.LookupEnv(TERKINS_USER)
	sysEnvPass, _ = os.LookupEnv(TERKINS_PASS)
	sysEnvEncrypted, _ = os.LookupEnv(TERKINS_ENCRYPTED)
	if envHost == "" {
		envHost = sysEnvHost
	}
	if envUser == "" {
		envUser = sysEnvUser
	}
	if envPass == "" {
		envPass = sysEnvPass
	}
	if envEncrypted == "" {
		envEncrypted = sysEnvEncrypted
	}
	isEncrypted = goter.IsYes(envEncrypted, true)
	if isEncrypted {
		hexBytes, _ := hex.DecodeString(envPass)
		envPass = string(goter.AesDecryptCBC(hexBytes, goter.GetKey(envUser)))
	}
}

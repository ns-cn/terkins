package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/ns-cn/goter"
	"os"
)

var session *gojenkins.Jenkins

func getSession() *gojenkins.Jenkins {
	if session != nil {
		return session
	}
	if goter.IsYes(envDebug, false) {
		fmt.Printf("terkins://%s@%s(using password: %s)\n", envHost, envUser, envPass)
	}
	jenkins := gojenkins.CreateJenkins(nil, envHost, envUser, envPass)
	_, err := jenkins.Init(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	session = jenkins
	return jenkins
}

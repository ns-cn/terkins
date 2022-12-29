package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"os"
	"terkins/env"
)

var session *gojenkins.Jenkins

func getSession() *gojenkins.Jenkins {
	if session != nil {
		return session
	}
	if env.Debug.Value {
		fmt.Printf("terkins://%s@%s(using password: %s)\n", env.Host.Value, env.User.Value, env.Pass.Value)
	}
	jenkins := gojenkins.CreateJenkins(nil, env.Host.Value, env.User.Value, env.Pass.Value)
	_, err := jenkins.Init(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	session = jenkins
	return jenkins
}

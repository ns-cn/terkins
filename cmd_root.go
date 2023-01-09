package main

import (
	"github.com/ns-cn/goter"
	"terkins/env"
)

var root = goter.NewRootCmd("terkins", "tool to operate jenkins with terminal", env.VERSION)

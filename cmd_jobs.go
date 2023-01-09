package main

import (
	"github.com/spf13/cobra"
	"terkins/env"
)

var CmdJobs = root.NewSubCommand(&cobra.Command{
	Use:   "jobs",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}, &env.Host, &env.User, &env.Pass, &env.Encrypt, &env.Debug)

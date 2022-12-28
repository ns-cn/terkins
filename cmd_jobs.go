package main

import (
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
)

var CmdJobs = goter.Command{Cmd: &cobra.Command{
	Use:   "jobs",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}}

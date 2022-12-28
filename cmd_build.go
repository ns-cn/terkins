package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"sync"
	"time"
)

var CmdBuild = goter.Command{Cmd: &cobra.Command{
	Use:   "build",
	Short: "构建具体的任务，可选命令行流方式或参数方式",
	Run: func(cmd *cobra.Command, args []string) {
		ReadSetting()
		goter.Required(envHost, func(u string) bool { return u != "" }, "run without envHost", func() { _ = cmd.Help() })
		goter.Required(envUser, func(u string) bool { return u != "" }, "run without username", func() { _ = cmd.Help() })
		goter.Required(envPass, func(u string) bool { return u != "" }, "run without password", func() { _ = cmd.Help() })
		getSession()
		jobsToBuild := make([]string, 0)
		reader := bufio.NewReader(os.Stdin)
		if len(args) != 0 {
			for _, arg := range args {
				jobsToBuild = append(jobsToBuild, arg)
			}
		} else {
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				jobsToBuild = append(jobsToBuild, string(line))
			}
		}
		waitGroup := sync.WaitGroup{}
		results := make(chan string, 0)
		for _, job := range jobsToBuild {
			toBuild := false
			if goter.IsYes(envBuildInfo, true) {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("build %s(Y/N)?", job)
				choice, _ := reader.ReadString('\n')
				choice = choice[:len(choice)-1]
				toBuild = goter.IsYes(choice, false)
			} else {
				toBuild = true
			}
			if toBuild {
				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					subJobNames := strings.Split(job, "->")
					ctx := context.Background()
					for _, subJobName := range subJobNames {
						if strings.HasPrefix(subJobName, "/job/") {
							subJobName = subJobName[5:]
						}
						buildId, err := session.BuildJob(ctx, subJobName, map[string]string{})
						if err != nil {
							results <- fmt.Sprintf("[FAILED]: building %s with errors: %v", subJobName, err)
							return
						}
						fmt.Printf("new Building JOB: %s queueID: %d\n", subJobName, buildId)
						subJob, err := session.GetJob(ctx, subJobName)
						if err != nil {
							results <- fmt.Sprintf("[FAILED]: building %s with errors: %v", subJobName, err)
							return
						}
						// 权益之计，原项目通过job获取指定队列ID的Build对象时url拼错，导致方法无法使用，暂替换成此方法
						targetBuild, err := subJob.GetLastBuild(ctx)
						for i := 0; i < 60; i++ {
							if err != nil || targetBuild.Info().QueueID != buildId {
								time.Sleep(500 * time.Millisecond)
								targetBuild, err = subJob.GetLastBuild(ctx)
							} else {
								break
							}
						}
						if err != nil {
							results <- fmt.Sprintf("[FAILED]: building %s with errors: %v", subJobName, err)
							return
						}
						for {
							if targetBuild.IsRunning(ctx) {
								time.Sleep(500 * time.Millisecond)
							} else {
								break
							}
						}
						results <- fmt.Sprintf("[SUCCESS]: building %s [%d] successed", subJobName, buildId)
					}
				}()
			}
		}
		go func() {
			waitGroup.Wait()
			results <- "EXIT"
		}()
		for {
			select {
			case result := <-results:
				if result == "EXIT" {
					goto EXIT
				}
				fmt.Println(result)
			}
		}
	EXIT:
	},
}}

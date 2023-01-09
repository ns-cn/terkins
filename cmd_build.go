package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"sync"
	"terkins/env"
	"time"
)

var CmdBuild = root.NewSubCommand(&cobra.Command{
	Use:   "build",
	Short: "构建具体的任务，可选命令行流方式或参数方式",
	Run: func(cmd *cobra.Command, args []string) {
		env.ReadSetting()
		goter.Required(env.Host.Value, func(u string) bool { return u != "" }, "run without Host", func() { _ = cmd.Help() })
		goter.Required(env.User.Value, func(u string) bool { return u != "" }, "run without username", func() { _ = cmd.Help() })
		goter.Required(env.Pass.Value, func(u string) bool { return u != "" }, "run without password", func() { _ = cmd.Help() })
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
		chanResults := make(chan buildResult, 0)
		for _, job := range jobsToBuild {
			toBuild := false
			if env.ShowBuildInfo.Value {
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
					subJobNames := strings.Split(job, ">")
					ctx := context.Background()
					for _, subJobName := range subJobNames {
						if strings.HasPrefix(subJobName, "/job/") {
							subJobName = subJobName[5:]
						}
						buildId, err := session.BuildJob(ctx, subJobName, map[string]string{})
						if err != nil {
							chanResults <- buildFail(subJobName, buildId, -1, err.Error())
							return
						}
						subJob, err := session.GetJob(ctx, subJobName)
						if err != nil {
							chanResults <- buildFail(subJobName, buildId, -1, err.Error())
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
							chanResults <- buildFail(subJobName, buildId, -1, err.Error())
							return
						}
						for {
							if targetBuild.IsRunning(ctx) {
								time.Sleep(500 * time.Millisecond)
							} else {
								break
							}
						}
						chanResults <- buildSuccess(subJobName, buildId, -1)
					}
				}()
			}
		}
		go func() {
			waitGroup.Wait()
			chanResults <- EXIT
		}()
		results := make([]buildResult, 0)
		for {
			select {
			case result := <-chanResults:
				if result == EXIT {
					goto EXIT
				}
				results = append(results, result)
			}
		}
	EXIT:
		table, _ := gotable.Create("模块", "构建ID", "结果", "耗时", "异常信息")
		for _, result := range results {
			if result.success {
				_ = table.AddRow([]string{result.name, strconv.Itoa(int(result.buildId)), "成功", strconv.Itoa(int(result.tokeTime)), result.err})
			} else {
				_ = table.AddRow([]string{result.name, strconv.Itoa(int(result.buildId)), "失败", strconv.Itoa(int(result.tokeTime)), result.err})
			}
		}
		fmt.Print(table.String())
	},
}, &env.Host, &env.User, &env.Pass, &env.Encrypt, &env.Debug, &env.ShowBuildInfo)

type buildResult struct {
	success  bool
	name     string
	buildId  int64
	tokeTime int64
	// 失败
	err string
}

var EXIT = buildResult{success: true, name: "EXIT", buildId: -1, tokeTime: -1}

func buildSuccess(name string, buildId, tokeTime int64) buildResult {
	return buildResult{
		success:  true,
		name:     name,
		buildId:  buildId,
		tokeTime: tokeTime,
	}
}

func buildFail(name string, buildId, tokeTime int64, error string) buildResult {
	return buildResult{
		success:  false,
		name:     name,
		buildId:  buildId,
		err:      error,
		tokeTime: tokeTime,
	}
}

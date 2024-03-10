package main

import (
	"github.com/sirupsen/logrus"
	"go-docker/cgroups"
	"go-docker/cgroups/subsystem"
	"go-docker/container"
	"os"
	"strings"
)

func Run(cmdArray []string, tty bool, res *subsystem.ResourceConfig) {
	parentProcess, writePipe := container.NewParentProcess(tty)
	if parentProcess == nil {
		logrus.Errorf("failed to new parent process")
		return
	}
	if err := parentProcess.Start(); err != nil {
		logrus.Errorf("parent start faile , err:%v", err)
		return
	}
	//添加资源限制
	cGroupManager := cgroups.NewCgroupManage("go-docker")
	//删除资源限制
	defer cGroupManager.Destroy()
	//设置资源限制
	cGroupManager.Set(res)
	//将容器进程，添加到各个subsystem挂载对应的cgroup中
	cGroupManager.Apply(parentProcess.Process.Pid)

	sendInitCommand(cmdArray, writePipe)
	parentProcess.Wait()
}
func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	_, _ = writePipe.WriteString(command)
	_ = writePipe.Close()
}

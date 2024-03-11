package subsystem

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

//获取cgroup在文件系统中的绝对路径

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRootPath, err := findCgroupMountPoint(subsystem)
	if err != nil {
		logrus.Error("find cgroup mount point err : %s ", err.Error())
		return "", err
	}
	cgroupTotalPath := path.Join(cgroupRootPath, cgroupPath)
	_, err = os.Stat(cgroupTotalPath)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(cgroupTotalPath, 0755); err != nil {
			return "", err
		}
	}
	return cgroupTotalPath, nil
}

// 找到挂载subsystem 的 hierarchy cgroup根节点所在的目录
func findCgroupMountPoint(subsystem string) (string, error) {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem && len(fields) > 4 {
				return fields[4], nil
			}
		}
	}
	return "", scanner.Err()
}

package subsystem

/**
资源限制接口
*/

type ResourceConfig struct {
	//内存限制
	MemoryLimit string
	//cpu时间片权重
	CpuShare string
	//cpu核数
	CpuSet string
}

type Subsystem interface {
	// Name 返回subsystem名字，如cpu,memory
	Name() string
	//Set 设置cgroup在这个subsystem中的资源限制
	Set(cgroupPath string, res *ResourceConfig) error
	// Remove 移除这个cgroup的资源限制
	Remove(cgroupPath string) error
	// Apply 将某个进程添加到cgroup中
	Apply(cgroupPath string, pid int) error
}

var (
	Subsystems = []Subsystem{
		&MemorySubsystem{},
		&CpuSubsystem{},
		&CpuSetSubsystem{},
	}
)

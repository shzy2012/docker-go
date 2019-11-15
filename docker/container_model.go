package docker

import "time"

//Container 容器信息
type Container struct {
	ID              string          `json:"Id"`
	Names           []string        `json:"Names"`
	Image           string          `json:"Image"`
	ImageID         string          `json:"ImageID"`
	Command         string          `json:"Command"`
	Created         int             `json:"Created"`
	Ports           []Ports         `json:"Ports"`
	Labels          interface{}     `json:"Labels"`
	State           string          `json:"State"`
	Status          string          `json:"Status"`
	HostConfig      HostConfig      `json:"HostConfig"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
	Mounts          []Mounts        `json:"Mounts"`
}

//ContainerDetail 描述单个容器信息
type ContainerDetail struct {
	ID              string          `json:"Id"`
	Created         time.Time       `json:"Created"`
	Path            string          `json:"Path"`
	Args            []string        `json:"Args"`
	State           State           `json:"State"`
	Image           string          `json:"Image"`
	ResolvConfPath  string          `json:"ResolvConfPath"`
	HostnamePath    string          `json:"HostnamePath"`
	HostsPath       string          `json:"HostsPath"`
	LogPath         string          `json:"LogPath"`
	Name            string          `json:"Name"`
	RestartCount    int             `json:"RestartCount"`
	Driver          string          `json:"Driver"`
	Platform        string          `json:"Platform"`
	MountLabel      string          `json:"MountLabel"`
	ProcessLabel    string          `json:"ProcessLabel"`
	AppArmorProfile string          `json:"AppArmorProfile"`
	ExecIDs         interface{}     `json:"ExecIDs"`
	HostConfig      HostConfig      `json:"HostConfig"`
	GraphDriver     interface{}     `json:"GraphDriver"`
	Mounts          []Mounts        `json:"Mounts"`
	Config          Config          `json:"Config"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
}

//Ports 端口
type Ports struct {
	IP          string `json:"IP"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

//Bridge 桥接
type Bridge struct {
	IPAMConfig          interface{} `json:"IPAMConfig"`
	Links               interface{} `json:"Links"`
	Aliases             interface{} `json:"Aliases"`
	NetworkID           string      `json:"NetworkID"`
	EndpointID          string      `json:"EndpointID"`
	Gateway             string      `json:"Gateway"`
	IPAddress           string      `json:"IPAddress"`
	IPPrefixLen         int         `json:"IPPrefixLen"`
	IPv6Gateway         string      `json:"IPv6Gateway"`
	GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
	MacAddress          string      `json:"MacAddress"`
	DriverOpts          interface{} `json:"DriverOpts"`
}

//Networks 网络
type Networks struct {
	Bridge Bridge `json:"bridge"`
}

//Mounts 挂载点
type Mounts struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Driver      string `json:"Driver"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
	Propagation string `json:"Propagation"`
}

//State 状态
type State struct {
	Status     string    `json:"Status"`
	Running    bool      `json:"Running"`
	Paused     bool      `json:"Paused"`
	Restarting bool      `json:"Restarting"`
	OOMKilled  bool      `json:"OOMKilled"`
	Dead       bool      `json:"Dead"`
	Pid        int       `json:"Pid"`
	ExitCode   int       `json:"ExitCode"`
	Error      string    `json:"Error"`
	StartedAt  time.Time `json:"StartedAt"`
	FinishedAt time.Time `json:"FinishedAt"`
}

type LogConfig struct {
	Type   string `json:"Type"`
	Config Config `json:"Config"`
}

type RestartPolicy struct {
	Name              string `json:"Name"`
	MaximumRetryCount int    `json:"MaximumRetryCount"`
}

//HostConfig 主机网络类型
type HostConfig struct {
	Binds                interface{}   `json:"Binds"`
	ContainerIDFile      string        `json:"ContainerIDFile"`
	LogConfig            LogConfig     `json:"LogConfig"`
	NetworkMode          string        `json:"NetworkMode"`
	PortBindings         interface{}   `json:"PortBindings"`
	RestartPolicy        RestartPolicy `json:"RestartPolicy"`
	AutoRemove           bool          `json:"AutoRemove"`
	VolumeDriver         string        `json:"VolumeDriver"`
	VolumesFrom          interface{}   `json:"VolumesFrom"`
	CapAdd               interface{}   `json:"CapAdd"`
	CapDrop              interface{}   `json:"CapDrop"`
	Capabilities         interface{}   `json:"Capabilities"`
	DNS                  []interface{} `json:"Dns"`
	DNSOptions           []interface{} `json:"DnsOptions"`
	DNSSearch            []interface{} `json:"DnsSearch"`
	ExtraHosts           interface{}   `json:"ExtraHosts"`
	GroupAdd             interface{}   `json:"GroupAdd"`
	IpcMode              string        `json:"IpcMode"`
	Cgroup               string        `json:"Cgroup"`
	Links                interface{}   `json:"Links"`
	OomScoreAdj          int           `json:"OomScoreAdj"`
	PidMode              string        `json:"PidMode"`
	Privileged           bool          `json:"Privileged"`
	PublishAllPorts      bool          `json:"PublishAllPorts"`
	ReadonlyRootfs       bool          `json:"ReadonlyRootfs"`
	SecurityOpt          interface{}   `json:"SecurityOpt"`
	UTSMode              string        `json:"UTSMode"`
	UsernsMode           string        `json:"UsernsMode"`
	ShmSize              int           `json:"ShmSize"`
	Runtime              string        `json:"Runtime"`
	ConsoleSize          []int         `json:"ConsoleSize"`
	Isolation            string        `json:"Isolation"`
	CPUShares            int           `json:"CpuShares"`
	Memory               int           `json:"Memory"`
	NanoCpus             int           `json:"NanoCpus"`
	CgroupParent         string        `json:"CgroupParent"`
	BlkioWeight          int           `json:"BlkioWeight"`
	BlkioWeightDevice    []interface{} `json:"BlkioWeightDevice"`
	BlkioDeviceReadBps   interface{}   `json:"BlkioDeviceReadBps"`
	BlkioDeviceWriteBps  interface{}   `json:"BlkioDeviceWriteBps"`
	BlkioDeviceReadIOps  interface{}   `json:"BlkioDeviceReadIOps"`
	BlkioDeviceWriteIOps interface{}   `json:"BlkioDeviceWriteIOps"`
	CPUPeriod            int           `json:"CpuPeriod"`
	CPUQuota             int           `json:"CpuQuota"`
	CPURealtimePeriod    int           `json:"CpuRealtimePeriod"`
	CPURealtimeRuntime   int           `json:"CpuRealtimeRuntime"`
	CpusetCpus           string        `json:"CpusetCpus"`
	CpusetMems           string        `json:"CpusetMems"`
	Devices              []interface{} `json:"Devices"`
	DeviceCgroupRules    interface{}   `json:"DeviceCgroupRules"`
	DeviceRequests       interface{}   `json:"DeviceRequests"`
	KernelMemory         int           `json:"KernelMemory"`
	KernelMemoryTCP      int           `json:"KernelMemoryTCP"`
	MemoryReservation    int           `json:"MemoryReservation"`
	MemorySwap           int           `json:"MemorySwap"`
	MemorySwappiness     interface{}   `json:"MemorySwappiness"`
	OomKillDisable       bool          `json:"OomKillDisable"`
	PidsLimit            int           `json:"PidsLimit"`
	Ulimits              interface{}   `json:"Ulimits"`
	CPUCount             int           `json:"CpuCount"`
	CPUPercent           int           `json:"CpuPercent"`
	IOMaximumIOps        int           `json:"IOMaximumIOps"`
	IOMaximumBandwidth   int           `json:"IOMaximumBandwidth"`
	MaskedPaths          []string      `json:"MaskedPaths"`
	ReadonlyPaths        []string      `json:"ReadonlyPaths"`
}

//Config 配置
type Config struct {
	Hostname     string      `json:"Hostname"`
	Domainname   string      `json:"Domainname"`
	User         string      `json:"User"`
	AttachStdin  bool        `json:"AttachStdin"`
	AttachStdout bool        `json:"AttachStdout"`
	AttachStderr bool        `json:"AttachStderr"`
	ExposedPorts interface{} `json:"ExposedPorts"`
	Tty          bool        `json:"Tty"`
	OpenStdin    bool        `json:"OpenStdin"`
	StdinOnce    bool        `json:"StdinOnce"`
	Env          []string    `json:"Env"`
	Cmd          []string    `json:"Cmd"`
	ArgsEscaped  bool        `json:"ArgsEscaped"`
	Image        string      `json:"Image"`
	Volumes      interface{} `json:"Volumes"`
	WorkingDir   string      `json:"WorkingDir"`
	Entrypoint   []string    `json:"Entrypoint"`
	OnBuild      interface{} `json:"OnBuild"`
	Labels       interface{} `json:"Labels"`
}

//NetworkSettings 网络设置
type NetworkSettings struct {
	Bridge                 string      `json:"Bridge"`
	SandboxID              string      `json:"SandboxID"`
	HairpinMode            bool        `json:"HairpinMode"`
	LinkLocalIPv6Address   string      `json:"LinkLocalIPv6Address"`
	LinkLocalIPv6PrefixLen int         `json:"LinkLocalIPv6PrefixLen"`
	Ports                  Ports       `json:"Ports"`
	SandboxKey             string      `json:"SandboxKey"`
	SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses"`
	SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
	EndpointID             string      `json:"EndpointID"`
	Gateway                string      `json:"Gateway"`
	GlobalIPv6Address      string      `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen    int         `json:"GlobalIPv6PrefixLen"`
	IPAddress              string      `json:"IPAddress"`
	IPPrefixLen            int         `json:"IPPrefixLen"`
	IPv6Gateway            string      `json:"IPv6Gateway"`
	MacAddress             string      `json:"MacAddress"`
	Networks               Networks    `json:"Networks"`
}

//GetID 容器ID
func (c *Container) GetID() string {
	return c.ID[0:12]
}

//GetImageID Image ID
func (c *Container) GetImageID() string {
	return c.ImageID[7:19]
}

package system

type Stats struct {
	Uptime int64
	Disks  []Disk
}

type Disk struct {
	MountPoint string
}

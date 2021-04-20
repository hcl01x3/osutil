package osutil

import (
	vers "github.com/hashicorp/go-version"
)

const (
	Unknown = "unknown"
	Ubuntu  = "ubuntu"
	Centos  = "centos"
	Redhat  = "redhat"
)

type Stats interface {
	Stats() map[string]float64
}

type Package struct {
	Name    string
	Version vers.Version
}

type PackageManager interface {
	Install(packs ...Package) error
	Uninstall(packs ...Package) error
	Update(packs ...Package) error
	IsInstalled(pack Package) error

	AddRepo(repos ...string) error
	RemoveRepo(repos ...string) error
	RefreshRepo() error
	CleanRepoCache()
}

type DistroManager interface {
	DistroName() string
	CodeName() string
	Version() *vers.Version
	CompareVersion(other *vers.Version) int
}

type CpuInfo interface {
	Stats

	Htz() uint
	NumCores() uint
	NumVCores() uint
}

type CpuManager interface {
	Cpus() []CpuInfo
	Cpu(num uint) CpuInfo
}

type MemManager interface {
	Size() int64
	Usage() int64
}

type DiskInfo interface {
	Stats

	Path() string
	MountPath() string
	Size() int64
	Usage() int64
	GetSubDisks() []DiskInfo
}

type DiskManager interface {
	Disks() ([]DiskInfo, error)
	DiskByPath(path string) (DiskInfo, error)
}

type NetDevInfo interface {
	Stats
}

type NetDevManager interface {
}

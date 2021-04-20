package osutil

import (
	"regexp"
)

/*
func NewPackageManager() (PackageManager, error) {
	return newPackageManager()
}
*/

func NewDistroManager() (DistroManager, error) {
	return newDistroManager()
}

/*
func NewCpuManager() (CpuManager, error) {
	return newCpuManager()
}

func NewMemManager() (MemManager, error) {
	return newMemManager()
}
*/

func NewDiskManager(match ...*regexp.Regexp) (DiskManager, error) {
	return newDiskManager(match...)
}

/*
func NewNetDevManger() (NetDevManager, error) {
	return NewNetDevManager()
}
*/

// +build linux

package osutil

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var _ DiskManager = (*diskManager)(nil)

var (
	diskPath  = "/sys/block"
	diskRegex = regexp.MustCompile(`^((sd|hd).+\d+)$`)
)

type disk struct {
	path      string
	mountPath string
	size      int64
	usage     int64
}

type diskManager struct {
	diskPath  string
	diskRegex *regexp.Regexp
	disks     []DiskInfo
}

func newDiskManager(match ...*regexp.Regexp) (*diskManager, error) {
	d := &diskManager{
		diskPath:  diskPath,
		diskRegex: diskRegex,
	}
	if len(match) > 0 && match[0] != nil {
		d.diskRegex = match[0]
	}
	return d, nil
}

func (d *diskManager) Disks() ([]DiskInfo, error) {
	files, err := ioutil.ReadDir(d.diskPath)
	if err != nil {
		return nil, err
	}
	disks := []DiskInfo{}
	for _, f := range files {
		if d.diskRegex.Match([]byte(f.Name())) {
			fmt.Println(f.Name())
		}
	}
	return disks, nil
}

func (d *diskManager) DiskByPath(path string) (DiskInfo, error) {
	return nil, nil
}

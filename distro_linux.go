// +build linux

package osutil

import (
	"os/exec"
	"strings"

	vers "github.com/hashicorp/go-version"
)

var _ DistroManager = (*distroManager)(nil)

var (
	distroCmd = "lsb_release"
	distroArg = "-irc"
)

type distroManager struct {
	distroName string
	codeName   string
	version    *vers.Version
}

func newDistroManager() (*distroManager, error) {
	d := &distroManager{
		distroName: Unknown,
		codeName:   Unknown,
		version:    vers.Must(vers.NewSemver("0.0.0")),
	}
	info, err := getDistroInfo()
	if err != nil {
		return d, nil
	}
	d.parseDistroInfo(info)
	return d, nil
}

func getDistroInfo() (string, error) {
	cmd := exec.Command(distroCmd, distroArg)
	if err := cmd.Run(); err != nil {
		return "", err
	}
	buf := []byte{}
	cmd.Stdin.Read(buf)
	return string(buf), nil
}

func (d *distroManager) parseDistroInfo(info string) error {
	for _, line := range strings.Split(info, "\n") {
		tokens := strings.SplitN(line, ":", 1)

		var key, val string

		if len(tokens) > 0 {
			key = strings.TrimSpace(tokens[0])
		}
		if len(tokens) > 1 {
			val = strings.TrimSpace(tokens[1])
		}

		switch key {
		case "Distributor ID":
			d.distroName = val
		case "Release":
			d.version = vers.Must(vers.NewSemver(val))
		case "Codename":
			d.codeName = val
		}
	}
	return nil
}

func (d distroManager) DistroName() string {
	return d.distroName
}

func (d distroManager) Version() *vers.Version {
	return d.version
}

func (d distroManager) CodeName() string {
	return d.codeName
}

func (d distroManager) CompareVersion(other *vers.Version) int {
	return d.version.Compare(other)
}

package vmfusion

import (
	"fmt"
	"os"

	"github.com/drone-runners/drone-runner-aws/oshelp"
	"github.com/drone-runners/drone-runner-aws/types"
	"github.com/sirupsen/logrus"
)

type Option func(*provider)

func SetPlatformDefaults(platform *types.Platform) (*types.Platform, error) {
	if platform.Arch == "" {
		platform.Arch = oshelp.ArchAMD64
	}
	// verify that we are using sane values for OS
	if platform.OS == "" {
		platform.OS = oshelp.OSMac
	}
	if platform.OS != oshelp.OSMac {
		return platform, fmt.Errorf("invalid OS %s, has to be '%s'", platform.OS, oshelp.OSMac)
	}

	return platform, nil
}

func WithUsername(username string) Option {
	return func(p *provider) {
		p.username = username
	}
}

func WithPassword(password string) Option {
	return func(p *provider) {
		p.password = password
	}
}

func WithISO(iso string) Option {
	return func(p *provider) {
		p.ISO = iso
	}
}

func WithCPU(cpu int64) Option {
	return func(p *provider) {
		p.CPU = cpu
	}
}

func WithMemory(memory int64) Option {
	return func(p *provider) {
		p.Memory = memory
	}
}

func WithVDiskPath(vDiskPath string) Option {
	return func(p *provider) {
		p.VDiskPath = vDiskPath
	}
}

// WithUserData returns an option to set the cloud-init template from a file location or passed in text.
func WithUserData(text, path string) Option {
	if text != "" {
		return func(p *provider) {
			p.userData = text
		}
	}
	return func(p *provider) {
		if path != "" {
			data, err := os.ReadFile(path)
			if err != nil {
				logrus.WithError(err).
					Fatalln("failed to read user_data file")
				return
			}
			p.userData = string(data)
		}
	}
}

// WithStorePath the path where VM machines are stored
func WithStorePath(storePath string) Option {
	return func(p *provider) {
		p.StorePath = storePath
	}
}

// WithRootDirectory sets the root directory for the virtual machine.
func WithRootDirectory(dir string) Option {
	return func(p *provider) {
		p.rootDir = tempdir(dir)
	}
}

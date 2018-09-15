package storage

import (
	"os"
	"fmt"
)

type Config struct {
	MountPath string // Directory where server mounts
}

func (c *Config) Validate() (error) {
	var (
		s   os.FileInfo
		err error
	)
	if s, err = os.Stat(c.MountPath); err != nil {
		return fmt.Errorf("storage config: %s, error: %s", c.MountPath, err)
	}
	if s.IsDir() == false {
		return fmt.Errorf("storage config: %s is not a directory", c.MountPath)
	}
	return nil
}

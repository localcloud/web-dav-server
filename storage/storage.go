package storage

import (
	"golang.org/x/net/webdav"
	"fmt"
	"os"
	"golang.org/x/net/context"
)

type storage struct {
	cfg     *Config
	tmpStub webdav.FileSystem
}

func (c *storage) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	return c.tmpStub.Mkdir(ctx, name, perm)
}

func (c *storage) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	return c.tmpStub.OpenFile(ctx, name, flag, perm)
}

func (c *storage) RemoveAll(ctx context.Context, name string) error {
	return c.tmpStub.RemoveAll(ctx, name)
}

func (c *storage) Rename(ctx context.Context, oldName, newName string) error {
	return c.tmpStub.Rename(ctx, oldName, newName)
}

func (c *storage) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	return c.tmpStub.Stat(ctx, name)
}

func storageLayerInitor(storage *storage) *storage {
	storage.tmpStub = webdav.Dir(storage.cfg.MountPath)
	return storage
}

// Fs Layer constructor
func New(cfg *Config) (webdav.FileSystem, error) {
	var (
		err error
	)
	if err = cfg.Validate(); err != nil {
		return nil, fmt.Errorf("storage layer: could not to create instance of storage, validation fails with: %s", err)
	}
	return storageLayerInitor(&storage{cfg: cfg}), nil
}

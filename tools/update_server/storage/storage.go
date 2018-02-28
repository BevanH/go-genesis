//go:generate sh -c "mockery -inpkg -name Engine -print > file.tmp && mv file.tmp storage_mock.go"
package storage

import "github.com/GenesisKernel/go-genesis/tools/update_server/model"

type Engine interface {
	GetVersionsList() ([]model.Version, error)
	Get(binary model.Build) (model.Build, error)
	Add(build model.Build) error
	Delete(binary model.Build) error
}

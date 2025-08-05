package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/repository"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/repository/memstorage"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/serverconfig"
)

type repositories struct {
	MemStorage repository.MemStorage
}

func NewRepositories(c *serverconfig.Config) (*repositories, error) {
	memStorage, err := memstorage.NewMemStorage(c.FileName, c.RestoreFileData, c.StoreInterval)
	if err != nil {
		return nil, err
	}

	return &repositories{
		MemStorage: memStorage,
	}, nil
}

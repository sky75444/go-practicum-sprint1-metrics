package app

import (
	"github.com/sky75444/go-practicum-sprint1-metrics.git/internal/repository"
	"github.com/sky75444/go-practicum-sprint1-metrics.git/internal/repository/memstorage"
)

type repositories struct {
	MemStorage repository.MemStorage
}

func NewRepositories() *repositories {
	return &repositories{
		MemStorage: memstorage.NewMemStorage(),
	}
}

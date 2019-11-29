package database

import (
	"go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
)

type PlatformRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *PlatformRepository) Create(int) (entities.Platform, error) {
	var platform entities.Platform
	platform.Id = 2
	platform.Platform = 2
	return platform, nil
}

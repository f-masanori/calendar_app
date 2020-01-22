package database

import (
	"go_docker/calender/entities"
	"go_docker/calender/infrastructure/database"
)

type PlatformRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *PlatformRepository) Create(int) (entities.Platform, error) {
	var platform entities.Platform
	platform.ID = 2
	platform.Platform = 2
	return platform, nil
}

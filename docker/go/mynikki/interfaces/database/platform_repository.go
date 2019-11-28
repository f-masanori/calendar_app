package database

import (
	"github.com/f-masanori/my-nikki_dev/docker/go/entities"
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
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

package services

import (
	"fmt"

	"go_docker/mynikki/entities"
)

type PlatformRepository interface {
	// FindAll() (entities.Users, error)
	// SaveUser(string)
	Create(int) (entities.Platform, error)
	// Find(int) (entities.User, error)
	// Save(*entities.User) (entities.User, error)
	// Update(int, *entities.User) (entities.User, error)
	// Delete(int) (error)
}
type PlatformService struct {
	PlatformRepository PlatformRepository
}

func (p *PlatformService) SavePlatform() {
	platform, err := p.PlatformRepository.Create(3)
	fmt.Println(platform)
	fmt.Println(err)
	fmt.Println("plat")
}

// Index
// func (d * DeviceService) GetAll() (entities.Users, error) {
// 	users, err := s.UserRepository.FindAll()
// 	return users, err
// }

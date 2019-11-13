package services

import (
	"github.com/f-masanori/my-nikki_dev/docker/go/entities"
)

type UserRepository interface {
	FindAll() (entities.Users, error)
	// Find(int) (entities.User, error)
	// Save(*entities.User) (entities.User, error)
	// Update(int, *entities.User) (entities.User, error)
	// Delete(int) (error)
}
type UserService struct {
	UserRepository UserRepository
}

// Index
func (s *UserService) GetAll() (entities.Users, error) {
	users, err := s.UserRepository.FindAll()
	return users, err
}

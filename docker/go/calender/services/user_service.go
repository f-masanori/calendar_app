package services

import (
	"fmt"
	"go_docker/calender/entities"
	"go_docker/calender/infrastructure/database"
	sqlcmd "go_docker/calender/interfaces/database"
	"strconv"
)

type UserRepository interface {
	FindAll() (entities.Users, error)
	CreateUser(string) (entities.User, error)
	DeleteUser(int) (int, error)
	// Find(int) (entities.User, error)
	// Save(*entities.User) (entities.User, error)
	// Update(int, *entities.User) (entities.User, error)
}
type UserService struct {
	UserRepository UserRepository
}

/* for test */
func NewUserService(sqlHandler *database.SqlHandler) *UserService {
	return &UserService{
		UserRepository: &sqlcmd.UserRepository{
			SqlHandler: sqlHandler,
		},
	}
}

/* ******** */

// Index
func (s *UserService) GetAll() (entities.Users, error) {
	users, err := s.UserRepository.FindAll()
	return users, err
}
func (s *UserService) StoreNewUser(name string) (entities.User, error) {
	fmt.Println("StoreNewUser")
	user, err := s.UserRepository.CreateUser(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created New user ID=" + strconv.Itoa(user.ID) + " name=" + user.Name)
	}
	return user, err
}
func (s *UserService) DeleteUser(ID int) int {
	fmt.Println("DeleteUser")
	returnId, err := s.UserRepository.DeleteUser(ID)
	fmt.Println(err)
	return returnId
}

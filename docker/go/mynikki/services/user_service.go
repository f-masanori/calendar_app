package services

import (
	"fmt"
	"strconv"
	"go_docker/mynikki/entities"
)

type UserRepository interface {
	FindAll() (entities.Users, error)
	CreateUser(string) (entities.User, error)
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
func (s *UserService) StoreNewUser(name string) (entities.User,error){
	fmt.Println("StoreNewUser")
	user, err := s.UserRepository.CreateUser(name)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Created New user id="+strconv.Itoa(user.Id)+" name="+user.Name)
	}
	return user,err
}

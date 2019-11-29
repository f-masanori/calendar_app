package database

import (
	"go_docker/mynikki/infrastructure/database"
	"testing"
)

func TestFindAll(t *testing.T) {
	DBhandler := database.NewSqlHandler()
	UserRepository := NewUserRepository(DBhandler)
	user, err = UserRepository.CreateUser("nana")
	if user.Name != "nana" {
		t.Errorf("gotdwed")
	}
}

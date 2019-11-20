package database

import (
	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
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

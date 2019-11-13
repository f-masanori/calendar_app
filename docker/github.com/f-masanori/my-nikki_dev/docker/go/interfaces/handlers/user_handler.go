package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
	sqlcmd "github.com/f-masanori/my-nikki_dev/docker/go/interfaces/database"
	"github.com/f-masanori/my-nikki_dev/docker/go/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(sqlHandler *database.SqlHandler) *UserHandler {
	return &UserHandler{
		Service: &services.UserService{
			UserRepository: &sqlcmd.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler  Index")
	// 抽象的にGetALL
	users, error := h.Service.GetAll()

	fmt.Println(users)
	if error != nil {
		fmt.Println(error)
		return
	}

	// json変換
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

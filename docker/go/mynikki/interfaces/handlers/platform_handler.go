package handlers

import (
	"fmt"
	"net/http"

	"go_docker/mynikki/infrastructure/database"
	sqlcmd "go_docker/mynikki/interfaces/database"
	"go_docker/mynikki/services"
)

type PlatformHandler struct {
	Service *services.PlatformService
}

func NewPlatformHandler(sqlHandler *database.SqlHandler) *PlatformHandler {
	return &PlatformHandler{
		Service: &services.PlatformService{
			PlatformRepository: &sqlcmd.PlatformRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (h *PlatformHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(323)
}

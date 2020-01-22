package handlers

import (
	"fmt"
	"net/http"

	"go_docker/calender/infrastructure/database"
	sqlcmd "go_docker/calender/interfaces/database"
	"go_docker/calender/services"
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

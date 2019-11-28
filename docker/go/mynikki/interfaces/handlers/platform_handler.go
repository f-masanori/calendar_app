package handlers

import (
	"fmt"
	"net/http"

	"github.com/f-masanori/my-nikki_dev/docker/go/infrastructure/database"
	sqlcmd "github.com/f-masanori/my-nikki_dev/docker/go/interfaces/database"
	"github.com/f-masanori/my-nikki_dev/docker/go/services"
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

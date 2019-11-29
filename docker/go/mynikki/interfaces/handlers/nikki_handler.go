package handlers

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"

	// "go_docker/mynikki/entities"
	"go_docker/mynikki/infrastructure/database"
	sqlcmd "go_docker/mynikki/interfaces/database"
	"go_docker/mynikki/services"
)

type NikkiHandler struct {
	Service *services.NikkiService
}

func NewNikkiHandler(sqlHandler *database.SqlHandler) *NikkiHandler {
	return &NikkiHandler{
		Service: &services.NikkiService{
			NikkiRepository: &sqlcmd.NikkiRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (h *NikkiHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nikkihandler index")
	h.Service.GetAll()
	/* handler call service  */
	// users, error := h.Service.GetAll() //GetAllの返り値はエンティティのusersでいい？
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	/* ************ */
	
	/* presenter */
	// users構造体 → json変換
	// json_users, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(json_users)
	/* ********* */
}

// func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
// 	/* handler マッピング*/
// 	type Request struct {
// 		Uid  string `json:"uid"`
// 		Name string `json:"name"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	request := new(Request)
// 	err := decoder.Decode(&request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(request)
// 	/* ******* */

// 	/* handler service呼び出し */
// 	user,err :=h.Service.StoreNewUser(request.Name)
// 	if err != nil{
// 		fmt.Println(err)
// 	}else{
// 		fmt.Println("succused call Service.StoreNewUser")
// 	}
// 	/* ******* */

// 	fmt.Println(entities.Platform_map["ios"])

// 	/* Presenter */
// 	json_user,err := json.Marshal(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(json_user)
// 	/* ******* */
// }


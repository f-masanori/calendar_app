package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/f-masanori/my-nikki_dev/docker/go/entities"
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
	users, error := h.Service.GetAll() //GetAllの返り値はエンティティのusersでいい？

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

func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	// jsonのrequest をマッピングする構造体
	type Request struct {
		Uid  string `json:"uid"`
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	h.Service.StoreNewUser()
	fmt.Println(entities.Platform_map["ios"])
	// fmt.Println(r.Body)
	// var req request

	// err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)

	// var req request
	// err := decoder.Decode(&req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(req.uid)

	defer r.Body.Close()
	// log.Println(t.Test)

	//Requestから nameを受け取り 変換して
	// users, error := h.Service.SaveUser()
}

func (h *UserHandler) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler  Test")
	// 抽象的にGetALL
	// users, error := h.Service.GetAll()

	// fmt.Println(users)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	// // json変換
	// res, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(res)
	// reqres.W.Write([]byte("uuu"))
}

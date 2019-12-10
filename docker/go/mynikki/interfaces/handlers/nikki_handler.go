package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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
	
	/* handler call service  */
	nikkis, err := h.Service.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	/* ************ */
	
	/* presenter */
	// users構造体 → json変換
	json_nikkis, err := json.Marshal(nikkis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_nikkis)
	/* ********* */
}

func (h *NikkiHandler) CreateNikki(w http.ResponseWriter, r *http.Request) {
	/* handler マッピング*/
	type Request struct {
		UserId  int     `json:"UserId"`
		Date    int 	`json:"Date"`
		Content string  `json:"Content"`
		Title   string 	`json:"Title"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */

	/* handler service呼び出し */
	nikki,err :=h.Service.CreateNikki(request.UserId,request.Date,request.Title,request.Content)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("succused call Service.StoreNewUser")
	}
	/* ******* */

	/* Presenter */
	json_nikki,err := json.Marshal(nikki)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_nikki)
	/* ******* */
}
func (h *NikkiHandler) EditNikki(w http.ResponseWriter, r *http.Request) {
	/* handler マッピング*/
	type Request struct {
		UserId  int     `json:"UserId"`
		Date    int 	`json:"Date"`
		Content string  `json:"Content"`
		Title   string 	`json:"Title"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */
	h.Service.EditNikki(request.UserId,request.Date,request.Title,request.Content)
}
func (h *NikkiHandler) DeleteNikki(w http.ResponseWriter, r *http.Request){
	/* handler マッピング*/
	type Request struct {
		UserId  int     `json:"UserId"`
		Date    int 	`json:"Date"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */
	/* service 呼び出し */
	confirmDelete := h.Service.DeleteNikki(request.UserId,request.Date)
	fmt.Println(confirmDelete)
	/* ******* */
	/* Presenter */
	json_confirmDelete,err := json.Marshal(confirmDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_confirmDelete)
	/* ******* */
}


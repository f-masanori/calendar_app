package services

import (
	"fmt"
	"go_docker/mynikki/entities"
)

type NikkiRepository interface {
	FindAll() (entities.Nikkis, error)
	CreateNikki(int,int,string,string) (entities.Nikki, error)
	DeleteNikki(int, int) (int,int,int,error)
}
type NikkiService struct {
	NikkiRepository NikkiRepository
}
/* nikki delete時に使用 */
type ConfirmDelete struct {
		UserId 		int
		Date   		int
		RowsAffect 	int
		Err 		error
}
func (n *NikkiService) GetAll() (entities.Nikkis, error){
	
	nikkis, err := n.NikkiRepository.FindAll()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(nikkis)

	return nikkis,err
}

func (n *NikkiService) CreateNikki(UserId int,Date int,Title string,Content string)(entities.Nikki, error) {
	nikki, err := n.NikkiRepository.CreateNikki(UserId,Date,Title,Content)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(nikki)
	return nikki,err
}

func (n *NikkiService) DeleteNikki(UserId int,Date int) *ConfirmDelete { 
	
	ConfirmDelete := new(ConfirmDelete)
	ConfirmDelete.UserId,ConfirmDelete.Date,ConfirmDelete.RowsAffect,ConfirmDelete.Err = n.NikkiRepository.DeleteNikki(UserId,Date)
	if ConfirmDelete.Err != nil{
		fmt.Println(ConfirmDelete.Err)
	}
	return ConfirmDelete
}

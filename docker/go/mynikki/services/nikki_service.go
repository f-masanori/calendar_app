package services

import (
	"fmt"
	"go_docker/mynikki/entities"
)

type NikkiRepository interface {
	FindAll() (entities.Nikkis, error)
	GetNikki(int, int)
	CreateNikki(int, int, string, string, int) (entities.Nikki, error)
	DeleteNikki(int, int) (int, int, int, error)
	EditNikki(int, int, string, string)
	RegisterPhoto()
}
type NikkiService struct {
	NikkiRepository NikkiRepository
}

/* nikki delete時に使用 */
type ConfirmDelete struct {
	UserId     int
	Date       int
	RowsAffect int
	Err        error
}

func (n *NikkiService) GetAll() (entities.Nikkis, error) {

	nikkis, err := n.NikkiRepository.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nikkis)

	return nikkis, err
}

func (n *NikkiService) CreateNikki(UserId int, Date int, Title string, Content string) (entities.Nikki, error) {
	nikki, err := n.NikkiRepository.CreateNikki(UserId, Date, Title, Content, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nikki)
	return nikki, err
}

func (n *NikkiService) EditNikki(UserId int, Date int, Title string, Content string) {
	n.NikkiRepository.EditNikki(UserId, Date, Title, Content)
}

func (n *NikkiService) DeleteNikki(UserId int, Date int) *ConfirmDelete {

	ConfirmDelete := new(ConfirmDelete)
	ConfirmDelete.UserId, ConfirmDelete.Date, ConfirmDelete.RowsAffect, ConfirmDelete.Err = n.NikkiRepository.DeleteNikki(UserId, Date)
	if ConfirmDelete.Err != nil {
		fmt.Println(ConfirmDelete.Err)
	}
	return ConfirmDelete
}

func (n *NikkiService) GetNikki(UserId int, Date int) {
	n.NikkiRepository.GetNikki(1, 20191211)
}

func (n *NikkiService) RegisterPhoto(UserId int, Date int) {
	n.NikkiRepository.RegisterPhoto()
}

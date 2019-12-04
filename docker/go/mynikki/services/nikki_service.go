package services

import (
	"fmt"
	"go_docker/mynikki/entities"
)

type NikkiRepository interface {
	FindAll() (entities.Nikkis, error)
	CreateNikki(int,int,string,string) (entities.Nikki, error)
}
type NikkiService struct {
	NikkiRepository NikkiRepository
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

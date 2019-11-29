package services

import (
	"fmt"
	"go_docker/mynikki/entities"
)

type NikkiRepository interface {
	FindAll()
	CreateNikki() (entities.Nikki, error)
}
type NikkiService struct {
	NikkiRepository NikkiRepository
}

func (n *NikkiService) GetAll() {
	nikki, err := n.NikkiRepository.CreateNikki()
	if err != nil{
		fmt.Println(err)
	}
	// fmt.Println(platform)
	fmt.Println(nikki)
	n.NikkiRepository.FindAll()
}

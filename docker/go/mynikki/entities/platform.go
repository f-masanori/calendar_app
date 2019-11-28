package entities

var Platform_map map[string]int = map[string]int{
	"ios":     0,
	"android": 1,
	"web":     2}

type Platform struct {
	Id       int
	Platform int
}

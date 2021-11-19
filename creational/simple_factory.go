package creational

import "fmt"

type Human interface {
	laugh()
	cry()
	talk()
}

type YellowHuman struct {
}

func (y *YellowHuman) laugh(){
	fmt.Println("yellow human laugh")
}

func (y *YellowHuman) cry(){
	fmt.Println("yellow human cry")
}
func (y *YellowHuman) talk(){
	fmt.Println("yellow human talk")
}

type WhiteHuman struct {
}

func (y *WhiteHuman) laugh(){
	fmt.Println("white human laugh")
}

func (y *WhiteHuman) cry(){
	fmt.Println("white human cry")
}
func (y *WhiteHuman) talk(){
	fmt.Println("white human talk")
}

type BlackHuman struct {
}

func (y *BlackHuman) laugh(){
	fmt.Println("black human laugh")
}

func (y *BlackHuman) cry(){
	fmt.Println("black human cry")
}
func (y *BlackHuman) talk(){
	fmt.Println("black human talk")
}

type factory struct {

}

func (c *factory) createHuman(t int) Human{
	switch t{
	case 0:
		return new(YellowHuman)
	case 1:
		return new(WhiteHuman)
	case 2:
		return new(BlackHuman)
	}
	return nil
}

func testFactory(){
	f := &factory{}
	h := f.createHuman(0)
	h.laugh()
	h.cry()
	h.talk()
	h = f.createHuman(1)
	h.laugh()
	h.cry()
	h.talk()
	h = f.createHuman(2)
	h.laugh()
	h.cry()
	h.talk()
}



package main
import "fmt"

type Shape interface{
	Draw()
}

type Rectangle struct{
}
func (r Rectangle) Draw(){
	fmt.Println("Rectangle drawn")
}

type Circle struct{
}
func (c Circle) Draw(){
	fmt.Println("Circle drawn")
}


type GraphicsEditor struct{
}
func (g *GraphicsEditor) DrawShape(s Shape){
	s.Draw()
}

func main(){
	g1 := GraphicsEditor{}
	c1 := Circle{}
	r1 := Rectangle{}
	g1.DrawShape(c1)
	g1.DrawShape(r1)
}

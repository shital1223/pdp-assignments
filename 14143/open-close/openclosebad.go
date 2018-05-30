package main
import "fmt"

type Shape struct{
	shapeName string
}

type Circle struct{
	m Shape
}
func(c *Circle) SetType(){
	c.m.shapeName = "Circle"
}

type Rectangle struct{
	m Shape
}
func(r *Rectangle) SetType(){
	r.m.shapeName = "Rectangle"
}

type GraphicEditor struct{
}
func(g *GraphicEditor) DrawShape(s Shape){
  if (s.shapeName=="Circle"){
	g.DrawCircle(s)
  }
  if(s.shapeName=="Rectangle"){
	g.DrawRectangle(s)
  }
}

func (g *GraphicEditor) DrawCircle(s Shape){
	fmt.Println(s.shapeName+" Drawn")
}

func (g *GraphicEditor) DrawRectangle(s Shape){
	fmt.Println(s.shapeName+" Drawn")
}

func main(){
  g := GraphicEditor{}
  r := Rectangle{}
  c := Circle{}
  c.SetType()
  r.SetType()
  g.DrawShape(r.m)
  g.DrawShape(c.m)
}


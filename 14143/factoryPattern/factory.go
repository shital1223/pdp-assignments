package main
import "fmt"

type Shape interface{
	Draw()
}


type Point struct {
	X,Y int
}

type Circle struct{
	Location Point
	Radius int
}
type Rectangle struct{
	Location Point
	Height int
	Width  int
}
type Square struct{
	Location Point
	Height int
	Width  int
}

func (c Circle) Draw() {
	fmt.Println("Circle Drawn")
}
func (r Rectangle) Draw() {
	fmt.Println("Rectangle Drawn")
}
func (s Square) Draw() {
	fmt.Println("Square Drawn")
}

type Factory struct{
//	obj Shape
}

func (f *Factory) createSquare(loc Point,w int,h int)(s Square){
		obj := Square{Location:loc,Height:h,Width:w}
		return  obj
}

func (f *Factory) createRectangle(loc Point,w int,h int)(r Rectangle){
		obj := Rectangle{Location:loc,Height:h,Width:w}
		return  obj
}

func (f *Factory) createCircle(loc Point,r int)(c Circle){
		obj := Circle{Location:loc,Radius:r}
		return  obj
}

func (f *Factory) GetType(s string,loc Point,w int,h int,r int)(shape Shape){
	if(s == "square"){
		return f.createSquare(loc,w,h)
	}
	if(s == "rectangle"){
		return f.createRectangle(loc,w,h)
	}
	if(s == "circle"){
		return f.createCircle(loc,r)
	}
	return
}

func main(){
	f := Factory{}
	l := Point{X:10,Y:10}
	cirobj := f.GetType("circle",l,10,30,4)
	cirobj.Draw()
	sqobj := f.GetType("square",l,10,30,4)
	sqobj.Draw()
	recobj := f.GetType("rectangle",l,10,30,4)
	recobj.Draw()
}

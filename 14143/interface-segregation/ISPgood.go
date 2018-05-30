package main
import "fmt"

type Workable interface{
	Work()
}

type Feedable interface{
	Eat()
}

type Worker struct{
}
func (w Worker) Work(){
	fmt.Println("In Worker's work method")
}
func (w Worker) Eat(){
	fmt.Println("In Worker's eat method")
}


type Robot struct{
}
func (r Robot) Work(){
	fmt.Println("In robot work method")
}


type Manager struct{
	iworker Workable
	feed Feedable
}
func (m *Manager) SetWorker(worker Workable){
	m.iworker = worker
}
func (m *Manager) Eat(f1 Feedable){
	m.feed = f1
}
func (m *Manager) Manage(){
	m.iworker.Work()
	m.feed.Eat()
}

func main(){
	w1 := Worker{}
	r1 := Robot{}
	m1 := Manager{}
	m1.SetWorker(w1)
	m1.Eat(w1)
	m1.Manage()
	m1.SetWorker(r1)
	m1.Manage()
}

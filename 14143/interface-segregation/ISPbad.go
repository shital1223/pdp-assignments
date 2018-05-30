package main
import "fmt"

type Iworker interface{
	Work()
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

type SuperWorker struct{
}
func (s SuperWorker) Work(){
	fmt.Println("In SuperWorker's work method")
}
func (s SuperWorker) Eat(){
	fmt.Println("In SuperWorker's eat method")
}

type Manager struct{
	worker Iworker
}
func (m *Manager) SetWorker(w Iworker){
	m.worker = w
}
func (m *Manager) Manage(){
	m.worker.Work()
}

func main(){
	w1 := Worker{}
	s1 := SuperWorker{}
	m1 := Manager{}
	m1.SetWorker(w1)
	m1.Manage()
	m1.SetWorker(s1)
	m1.Manage()
}

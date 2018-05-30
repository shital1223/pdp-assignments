package main
import "fmt"

type Iworker interface{
	Work()
}

type Worker struct{}
func (w Worker) Work() {
	fmt.Println("In Worker's work method")
}

type SuperWorker struct{}
func (s SuperWorker) Work(){
	fmt.Println("In SuperWorker's work method")
}

type Manager struct{
	iworker Iworker
}
func(m *Manager) SetWorker(iw Iworker){
	m.iworker = iw
}
func(m *Manager) Manage(){
	m.iworker.Work()
}

func main(){
	w := Worker{}
	s := SuperWorker{}
	m := Manager{}
	m.SetWorker(w)
	m.Manage()
	m.SetWorker(s)
	m.Manage()
}

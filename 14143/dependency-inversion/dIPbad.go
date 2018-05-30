package main
import "fmt"

type Worker struct{}
func  (w *Worker) Work(){
	fmt.Println("In Worker's work method")
}

type  SuperWorker struct{}
func (s *SuperWorker) Work(){
	fmt.Println("In SuperWorker's method")
}

type Manager struct{
	worker Worker
	superworker SuperWorker
}
func (m *Manager) SetWorker(w Worker){
		m.worker = w
}
func (m *Manager) SetSuperWorker(s SuperWorker){
		m.superworker = s
}
func (m *Manager) Manage(){
		m.worker.Work()
		m.superworker.Work()
}

func main(){
	w1 := Worker{}
	s1 := SuperWorker{}
	m1 := &Manager{}
	m1.SetWorker(w1)
	m1.SetSuperWorker(s1)
	m1.Manage()
}

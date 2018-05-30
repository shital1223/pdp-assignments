package main
import "fmt"


type Customer struct {
	custID string
	subject string
	body string
}

type EmailSending struct{
}

func (e *EmailSending) GetCustomer(ID string,sub string,msg string) Customer {
	cust := Customer{custID:ID,subject:sub,body:msg}
	return cust
}
func (e *EmailSending) GetContent(cust Customer) string{
	var emailContent =cust.custID+" "+cust.subject+" "+cust.body
	return emailContent
}
func (e *EmailSending) SendMail(msg string){
	fmt.Println("message is sent to :"+msg)
}

func main(){
	e1 := EmailSending{}
	cust := e1.GetCustomer("shital@gmail.com","subject :bad example","body:single responsibilityprinciple")
	msg := e1.GetContent(cust)
	e1.SendMail(msg)
}

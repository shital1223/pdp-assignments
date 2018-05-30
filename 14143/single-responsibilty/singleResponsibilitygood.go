package main
import "fmt"

type Customer struct{
	custID string
       s subject string
	body string
}

type EmailSender struct{
}
func (s *EmailSender) SendMail(msg string){
	fmt.Println("Mail is sent to:"+msg)
}

type CustomerRepository struct{
}
func (c *CustomerRepository) GetCustomer(EID string,sub string,msg string) Customer{
	cust := Customer{custID:EID,subject:sub,body:msg}
	return cust
}

type EmailContentBuilder struct{
}
func (e *EmailContentBuilder) GetContent(cust Customer) string {
	var emailContent = cust.custID+" "+cust.subject+" "+cust.body
	return emailContent
}

func main(){
	c1 := &CustomerRepository{}
	cust := c1.GetCustomer("shital@gmail.com","PDP","single-responsibility principle")
	e1 := EmailContentBuilder{}
	msg := e1.GetContent(cust)
	s1 := EmailSender{}
	s1.SendMail(msg)
}

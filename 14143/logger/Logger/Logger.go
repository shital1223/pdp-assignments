package Logger
import (
	"fmt"
	"os"
	"io"
	"time"
	)
var FileName string
var LevelID int
var LevelName string

func Init(fileName string,levelName string){
	FileName = fileName 
	LevelName = levelName
	SetLevelID(LevelName)
	f, err := os.Create(FileName)
	if err !=nil{
		fmt.Println(err)
	}
	f.Close()
}

func SetLevelID(levelname string) {
	switch levelname {
		case "DEBUG":
			LevelID = 1
		case "INFO":
			LevelID = 2
		case "WARNING":
			LevelID = 3
		case "ERROR":
			LevelID = 4
		case "FATAL":
			LevelID = 5
	}
}

func WriteLog(message string){
	f, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil{
		panic(err)
	}
	io.WriteString(f,message)
	defer f.Close()
}

func DEBUG(message string){
	t := (time.Now().Format(time.RFC850))
	msg := "DEBUG : "+t+" : "+message
	if(LevelID<=1){
		WriteLog(msg)
	}
}

func INFO(message string){
	t := (time.Now().Format(time.RFC850))
	msg := "INFO : "+t+" : "+message
	if(LevelID<=2){
		WriteLog(msg)
	}
}

func WARNING(message string){
	t := (time.Now().Format(time.RFC850))
	msg := "WARNING : "+t+" : "+message
	if(LevelID<=3){
		WriteLog(msg)
	}
}

func ERROR(message string){
	t := (time.Now().Format(time.RFC850))
	msg := "ERROR : "+t+" : "+message
	if(LevelID<=4){
		WriteLog(msg)
	}
}

func FATAL(message string){
	t := (time.Now().Format(time.RFC850))
	msg := "FATAL : "+t+" : "+message
	if(LevelID<=5){
		WriteLog(msg)
	}
}
/*
func main(){
	l1 := GetInstance("sample.log","WARNING");
	// check if multiple instace created or not.
	//l2 := GetInstance("sam.l","DEBUG");
	fmt.Println(l1)
	fmt.Println(l2)
	l1.CreateFile("sample.log")
	message := "this is the message\n"
	l1.INFO(message)
	l1.WARNING(message)
	l1.ERROR(message)
	l1.FATAL(message)
	l1.DEBUG(message)
}*/

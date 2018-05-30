package main
import (
	"Logger"
	)

func main(){
	Logger.Init("sample2.log","DEBUG")
   message1 := "this is debug message\n"
   message2 := "this is info message\n"
   message3 := "this is warning message\n"
   message4 := "this is error message\n"
   message5 := "this is fatal message\n"
   Logger.DEBUG(message1)
   Logger.INFO(message2)
   Logger.WARNING(message3)
   Logger.ERROR(message4)
   Logger.FATAL(message5)
}

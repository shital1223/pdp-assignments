package main
import (
	"fmt"
	"pool"
	_ "github.com/lib/pq"
	)

func main(){
	pool := &pool.ObjectPool{Size: 10}
   status := pool.Init()
	if nil != status {
		fmt.Println(status)
	}
	for i := 0; i < 10; i++ {
		object := pool.Borrow()
		fmt.Println("Object is acquired")
		pool.Release(object)
		fmt.Println("Object is release")
   }
}

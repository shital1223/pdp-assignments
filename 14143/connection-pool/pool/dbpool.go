package pool
import (
	"database/sql"
)

type ObjectPool struct{
  Size      int
  PoolChan  chan interface{}
}

func (o *ObjectPool) Init() {
	o.PoolChan = make(chan interface{},o.Size)
	for i := 0; i < o.Size; i++ {
		psqlInfo := "host='localhost' port=5432 user='postgres' password='2312' dbname='pdp' sslmode=disable"
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
		}
		o.PoolChan <- db
	}
}

func (o *ObjectPool) Borrow() interface{} {
   return <-o.PoolChan
}

func (o *ObjectPool) Release(object interface{}) {
	   o.PoolChan <- object
}



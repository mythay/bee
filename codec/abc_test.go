package codec

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_condition(t *testing.T) {
	con := sync.NewCond(new(sync.Mutex))
	v := 0
	go func() {
		time.Sleep(time.Second * 2)
		con.L.Lock()
		v = 1
		con.L.Unlock()
		con.Broadcast()
	}()
	con.L.Lock()
	for v == 0 {
		con.Wait()
	}

	con.L.Unlock()
	fmt.Println("done")
}

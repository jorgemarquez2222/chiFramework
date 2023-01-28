package fino

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func factorial(num int) int {
	if num <= 1 {
		return num
	}
	return num * factorial(num-1)
}

func Fact(proceso int, wg *sync.WaitGroup) {
	limit := 6
	wg.Add(limit)
	for i := 0; i < limit; i++ {
		go func(j int) {
			time.Sleep(time.Second * 1)
			fmt.Printf("proceso: %d i: %d valor: %d\n", proceso, j, factorial(j))
			wg.Done()
		}(i)
	}
}

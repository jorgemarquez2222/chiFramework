package fino

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

var mu sync.Mutex

func factorial(num int) int {
	if num <= 1 {
		return num
	}
	return num * factorial(num-1)
}

func Fact(proceso int) {
	mu.Lock()
	limit := 6
	var wgInterno sync.WaitGroup
	wgInterno.Add(limit)
	mu.Unlock()
	for i := 0; i < limit; i++ {
		go func(j int) {
			mu.Lock()
			time.Sleep(time.Second * 1)
			if factorial(j) > 5 {
				color.Cyan("proceso: %d i: %d valor: %d\n", proceso, j, factorial(j))
			} else {
				color.Red("proceso: %d i: %d valor: %d\n", proceso, j, factorial(j))
			}
			wgInterno.Done()
			mu.Unlock()
		}(i)
	}
	wgInterno.Wait()
}

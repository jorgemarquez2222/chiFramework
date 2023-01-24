package canales

import "fmt"

func MetodoSignal() {
	ch := make(chan string, 2) // un canal de 2 buffer
	go func() {
		fmt.Print(<-ch)
	}()
	send(ch)
}

func send(ch chan string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
}

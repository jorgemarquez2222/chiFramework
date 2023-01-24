package canales

import "fmt"

func MetodoSignal() {
	ch := make(chan string, 2)    // un canal de 2 buffer
	signal := make(chan struct{}) // un canal de 2 buffer
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		signal <- struct{}{}
	}()
	send(ch, signal)
	<-signal
}

func send(ch chan string, signal chan<- struct{}) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
}

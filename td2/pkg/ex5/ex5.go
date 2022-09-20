package ex5

import (
	"fmt"
	"sync"
	"time"
)

type Agent interface {
	Start()
}

type PingAgent struct {
	ID string
	C  chan string
}

func (pinga *PingAgent) Start(wg *sync.WaitGroup, pongchan* chan string) {
	defer wg.Done()
	for{
	time.Sleep(time.Second)
	select {
	case c := <-pinga.C:
		if c == "Pong !" {
			*pongchan <- "Ping !"
			fmt.Println("in Ping: Pong received !")
		} else {
			fmt.Println("in Ping: Unknown message !")
		}
	case <- time.After(3 * time.Second):
		fmt.Println("in Ping: timed out : no signal received")
		*pongchan <- "Ping !"
	}
	}
}

type PongAgent struct {
	ID string
	C  chan string
}

func (ponga *PongAgent) Start(wg *sync.WaitGroup, pingchan* chan string) {
	defer wg.Done()
	for{
	time.Sleep(time.Second)
	select {
	case c := <-ponga.C:
		if c == "Ping !" {
			*pingchan <- "Pong !"
			fmt.Println("in Pong: Ping received !")
		} else {
			fmt.Println("in Pong: Unknown message !")
		}
	case <- time.After(3 * time.Second):
		fmt.Println("in Pong: timed out : no signal received")
	}
	}
}


func RunPingPong() {
	var wg = new(sync.WaitGroup)
	wg.Add(2)
	// ch := make(chan string)
	fmt.Printf("DEBUG1")
	fmt.Printf("DEBUG2")
	pingagent := PingAgent{ID: "001", C: make(chan string)}
	fmt.Printf("DEBUG3")
	pongagent := PongAgent{ID: "002", C: make(chan string)}
	fmt.Printf("DEBUG4")
	go pingagent.Start(wg, &pongagent.C)
	fmt.Printf("DEBUG5")
	go pongagent.Start(wg, &pingagent.C)
	fmt.Printf("DEBUG6")
	wg.Wait()
}
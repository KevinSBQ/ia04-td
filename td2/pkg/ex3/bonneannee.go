package ex3

import(
	"fmt"
	"time"
)

func Ba1(){
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i!= 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Bonne année !")
}

/*
exemple of After: Timeout

select {
	case m := <-c:
		handle(m)
	case <-time.After(5 * time.Minute):
		fmt.Println("timed out")
}
*/

func Ba2(){
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i!= 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		// After wait for a period of time, pass the actual time into a channel and return this channel (<-chan Time)
		<-time.After(time.Second)
	}
	fmt.Println("Bonne année !")
}

func Ba3(){
	// Tick pass the actual time into a channel every period of time and return this channel (<-chan Time)
	c := time.Tick(time.Second) // attention aux fuite de memoire
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i!= 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		// reads from the channel: wait for the passing of time from Tick every second
		<-c
	}
	fmt.Println("Bonne année !")
}
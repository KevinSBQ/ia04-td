package ex3

import(
	"fmt"
	"time"
)

func ba1(){
	for i := 5; i > 0; i-- {
		if i!= 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Bonne année !")
}

func ba2(){
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i!= 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		time.After(time.Second)
	}
	fmt.Println("Bonne année !")
}

func ba3(){
	c := time.Tick(time.Second)
	for i := 5; i > 0; i-- {
		fmt.Print(i)
		if i!= 1 {
			fmt.Print(", ")
		} else {
			fmt.Print("... ")
		}
		<-c
	}
	fmt.Println("Bonne année !")
}
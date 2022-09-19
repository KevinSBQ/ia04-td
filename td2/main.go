package main

import (
	"fmt"
	"math"
	"td2/pkg/ex1"
	"td2/pkg/ex2"

	// "td2/pkg/ex3"
	"td2/pkg/ex4"
	"time"
)

const big = 5000000
const number_worker = 11

func main(){
/* EX1 */
	go ex1.CompteMsg(300, "Go 1")
	go ex1.CompteMsg(300, "Go 2")
	time.Sleep(100 * time.Millisecond)
	for i:=0; i < 10; i++ {
		go ex1.CompteMsgFromTo(i*10, (i+1)*10, fmt.Sprintf("Let's go %v", i))
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("\n\n\n")

/* EX2 */
	ex2.Ex2_original()
	ex2.Ex2_mutex()
	ex2.Ex2_waitgroup()

	fmt.Printf("\n\n\n")

/* EX4 */
	var tab1, tab2 [big]int
	var t1 time.Time
	for j := 0; j<10; j++ {
		fmt.Println("Fill Loop", j, "----------")
		t1 = time.Now()
		ex4.Fill(tab1[:], 10)
		fmt.Println("Time escaped : ", time.Since(t1))
		fmt.Println("FillConc Loop", j, "----------")
		t1 = time.Now()
		ex4.FillConc(tab1[:], 10, number_worker)
		fmt.Println("Time escaped conc: ", time.Since(t1))
	}


	for j:=0; j<10; j++ {
		fmt.Println("ForEach Loop", j, "----------")
		t1 = time.Now()
		ex4.ForEach(tab1[:], func(i int) int {return int(math.Sqrt(float64(i)))})
		fmt.Println("Time escaped : ", time.Since(t1))
		fmt.Println("ForEachConc Loop", j, "----------")
		t1 = time.Now()
		ex4.ForEachConc(tab1[:], func(i int) int {return int(math.Sqrt(float64(i)))}, number_worker)
		fmt.Println("Time escaped conc: ", time.Since(t1))
	}

	for j:=0; j<10; j++ {
		fmt.Println("Equal Loop", j, "----------")
		t1 = time.Now()
		ex4.Equal(tab1[:], tab2[:])
		fmt.Println("Time escaped : ", time.Since(t1))
		fmt.Println("EqualConc Loop", j, "----------")
		t1 = time.Now()
		ex4.EqualConc(tab1[:], tab2[:], number_worker)
		fmt.Println("Time escaped conc: ", time.Since(t1))
		fmt.Println("EqualConcInterup Loop", j, "----------")
		t1 = time.Now()
		ex4.EqualConcInterup(tab1[:], tab2[:], number_worker)
		fmt.Println("Time escaped conc: ", time.Since(t1))
	}

	fmt.Printf("\n\n\n")
}
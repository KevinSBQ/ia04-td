package main

import (
	"fmt"
	"math"
	"td2/pkg/ex1"
	"td2/pkg/ex2"
	// "td2/pkg/ex3"
	"td2/pkg/ex4"
	"td2/pkg/ex5"
	"time"
)

const big = 5000000
const number_worker = 11

func main(){
/* EX1 */
	fmt.Printf("\n1.1\n")
	ex1.Compte(100)
	go ex1.Compte(100)

	fmt.Printf("\n1.2\n")
	go ex1.CompteMsg(100, "Go 1")
	go ex1.CompteMsg(100, "Go 2")

	fmt.Printf("\n1.3 1.4\n")
	time.Sleep(100 * time.Millisecond)
	for i:=0; i < 10; i++ {
		go ex1.CompteMsgFromTo(i*10, (i+1)*10, fmt.Sprintf("Let's go %v", i))
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("\n\n")

/* EX2 */
	ex2.Ex2_original()
	ex2.Ex2_mutex()
	ex2.Ex2_waitgroup()

	fmt.Printf("\n\n")

/* EX3 */
	// ex3.Ba1()
	// ex3.Ba2()
	// ex3.Ba3()
	fmt.Printf("\n\n")

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

	ex4.Fill(tab1[:], 10)
	ex4.Fill(tab2[:], 10)

	for j:=0; j<10; j++ {
		fmt.Println("Equal Loop", j, "----------")
		t1 = time.Now()
		res1 := ex4.Equal(tab1[:], tab2[:])
		fmt.Println("Time escaped : ", time.Since(t1))
		fmt.Println("EqualConc Loop", j, "----------")
		t1 = time.Now()
		res2 := ex4.EqualConc(tab1[:], tab2[:], number_worker)
		fmt.Println("Time escaped conc: ", time.Since(t1))
		fmt.Println("EqualConcInterup Loop", j, "----------")
		t1 = time.Now()
		res3 := ex4.EqualConcInterup(tab1[:], tab2[:], number_worker)
		fmt.Println("Time escaped concinterup: ", time.Since(t1))

		fmt.Println(res1, res2, res3)
	}
	// Execution time needed: EqualConc < EqualConcInterup < Equal
	// Write an read from a channel is time-costly

	fmt.Printf("\n\n")

	// cannot read and write at same time
	// var cccc chan int
	// cccc <- 5
	// fmt.Println(<-cccc)

/* EX5 */
	ex5.RunPingPong()
}
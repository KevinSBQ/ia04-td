package pairimpair

import "fmt"

func PairImpair() {
	fmt.Println("Nombres pair 1-100")
	for i:=1; i<=100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Printf("\n")
	fmt.Println("Nombres impair 1-100")
	for i:=1; i<=100; i++ {
		if i%2 != 0 {
			fmt.Printf("%v ", i)
		}
	}
}
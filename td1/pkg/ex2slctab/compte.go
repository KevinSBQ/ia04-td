package ex2slctab

import "fmt"

func Compte(n int, tab []int){
	fmt.Printf("%v Numbers in tab : ", n) 
	for i:=0; i<n; i++ {
		fmt.Printf("%v ", tab[i])
	}
	fmt.Printf("\n")
}
package ex1

import (
	"fmt"
)

func Compte(n int) {
	for i:=0 ; i<n ; i++ {
		fmt.Println(i)
	}
}

func CompteMsg(n int, msg string){
	for i:=0; i<n; i++ {
		fmt.Println(msg, i)
	}
}

func CompteMsgFromTo(start int, end int, msg string){
	for i:= start; i<end; i++{
		fmt.Println(msg, i)
	}
}


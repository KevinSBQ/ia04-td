package main

import (
	"fmt"
	"td3/comsoc"
)

func main(){
	fmt.Println("Helloworld!")

	prefs := [][]comsoc.Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	fmt.Println(comsoc.STV_SCF(prefs))
}
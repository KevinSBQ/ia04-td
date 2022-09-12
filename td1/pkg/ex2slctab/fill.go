package ex2slctab

import (
	"math/rand"
)

func Fill(sl []int) {
	for i:=0; i<len(sl); i++ {
		sl[i] = rand.Intn(24) //returns an int in a half-open interval [0,n)
	}
}
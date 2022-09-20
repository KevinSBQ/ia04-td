package ex2

import (
    "fmt"
	"sync"
)

var n_mu = 0

func f1(wg *sync.WaitGroup) {
    defer wg.Done()
	n_mu++
}

// Waitgroup only ensures every goroutine ends before the end of function
func Ex2_waitgroup() {
	var wg = new(sync.WaitGroup)
	wg.Add(10000)
    for i := 0; i < 10000; i++ {
        go f1(wg)
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("WaitGroup : n_mu:", n_mu)
}
package ex2

import (
    "fmt"
	"sync"
)

type SynchronizedInt struct {
	sync.Mutex
	n int
}

var i SynchronizedInt

func f2() {
	i.Lock()
	defer i.Unlock()
	i.n++
}

func Ex2_mutex() {
    for i := 0; i < 10000; i++ {
        go f2()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("Mutex: i.n:", i.n)
}
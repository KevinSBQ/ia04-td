package ex2

import (
    "fmt"
)

var n = 0

func f() {
    n++
}

func Ex2_original() {
    for i := 0; i < 10000; i++ {
        go f()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("Original Version: n:", n)
}

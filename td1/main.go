package main

import (
	"fmt"
	"td1/bin/hello"
	"td1/bin/hi"
    "td1/bin/pairimpair"
	"td1/pkg/toto"
    "td1/pkg/ex2slctab"
)

func main() {
    fmt.Println("Hello, world.")
    PrintHi()
    toto.Toto1()
    toto.Toto2()
    toto.Toto3()
    hello.Hello()
    hi.Hi()

    pairimpair.PairImpair()

    var tab[4] int
    sl := tab[:] //definir un slice qui prend tous les element de tab
    ex2slctab.Fill(sl)
    fmt.Println("tab : ", tab, "sl : ", sl)
}
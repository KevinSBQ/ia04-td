package main

import (
	"fmt"
	"td1/bin/hello"
	"td1/bin/hi"
    "td1/bin/pairimpair"
	"td1/pkg/toto"
    "td1/pkg/ex2slctab"
    "td1/pkg/probleme"
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
    fmt.Println("Moyenne (use iteration): ", ex2slctab.Moyenne(sl))
    fmt.Println("Moyenne (use range of slice): ", ex2slctab.MoyenneRange(sl))
    fmt.Println("Moyenne (implicite return): ", ex2slctab.MoyenneImplicite(sl))


    fmt.Printf("\n")
    fmt.Println("Valeur Centrale(s) : ", ex2slctab.ValeurCentrales(sl))
    var tab2[5] int
    sl2 := tab2[:]
    ex2slctab.Fill(sl2)
    fmt.Println("tab2 : ", tab2, "sl2 : ", sl2)

    fmt.Printf("\n")
    ex2slctab.Plus1(sl2)
    fmt.Println("tab2 : ", tab2, "sl2 : ", sl2)

    fmt.Println("Valeur Centrale(s) : ", ex2slctab.ValeurCentrales(sl2))

    ex2slctab.Compte(5, sl2)

    fmt.Println(probleme.IsPalindrome("RADAR"))
    fmt.Println(probleme.IsPalindrome("RADAE"))
    dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
    dictsl := dict[:]
    fmt.Println(probleme.Palindromes(dictsl))
}
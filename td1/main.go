package main

import (
	"fmt"
	"sort"
	"td1/bin/hello"
	"td1/bin/hi"
	"td1/bin/pairimpair"
	"td1/pkg/ex2slctab"
	"td1/pkg/probleme"
	"td1/pkg/toto"
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
    fmt.Println(probleme.Footprint("AGENT"))
    fmt.Println(probleme.Anagrams(dictsl))

    dictfromfile := probleme.DictFromFile("dico-scrabble-fr.txt")
    // fmt.Println(dictfromfile)

    fmt.Println("Quel est le(s) plus long(s) palindromes de la langue française ?")
    palindromesfromfile := probleme.Palindromes(dictfromfile)
    sort.Slice(palindromesfromfile, func(i, j int) bool {
        return len(palindromesfromfile[i]) < len(palindromesfromfile[j])
    })
    fmt.Println(palindromesfromfile[len(palindromesfromfile)-1])
    anagramsfromfile := probleme.Anagrams(dictfromfile)

    fmt.Println("Quels sont les anagrammes de agents ?")
    fmt.Println(anagramsfromfile["AEGNT"])

    fmt.Println("Quel(s) mot(s) de la langue française contien(nen)t le plus d’anagrammes ?")
    values := make([][]string, 0, len(anagramsfromfile))
    for  _, value := range anagramsfromfile {
        values = append(values, value)
    }
    sort.Slice(values, func(i, j int) bool {
        return len(values[i]) < len(values[j])
    })
    if len(values)-1 > 0 {
        fmt.Println("Empreinte: ", probleme.Footprint(values[len(values)-1][0]))
        fmt.Println(values[len(values)-1])
    }
    

    fmt.Println("Existe-t-il un palindrome qui possède des anagrammes ?")
    keys := make([]string, 0, len(anagramsfromfile))
    for k := range anagramsfromfile {
        keys = append(keys, k)
    }
    // fmt.Println(keys)
    for _, p := range keys {
        if probleme.IsPalindrome(p) {
            // fmt.Println(p)
            fmt.Printf("%v ", p)
        }
    }
    fmt.Printf("\n")
    // Non, il n'y a pas
}
package main

import (
	"fmt"
	"td4/vote/restclientagent"
	"math/rand"
	rad "td4/vote"
)

func main() {
	nCandidat := 10
	permutation := rand.Perm(nCandidat)
		prefs := make([]rad.Alternative, 0, 10)
		for i := range permutation {
			permutation[i] += 1
			prefs[i] = rad.Alternative(permutation[i])
		}
	ag := restclientagent.NewRestClientAgent("id1", "http://localhost:8000", prefs)
	ag.Start()
	fmt.Scanln()
}

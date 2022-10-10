package main

import (
	"fmt"
	"log"
	"math/rand"
	"td4/vote/restclientagent"
	"td4/vote/restserveragent"
	rad "td4/vote"
)

func main() {
	const nAgent = 100
	const nCandidat = 10
	const url1 = ":8080"
	const url2 = "http://localhost:8080"
	// allocates an underlying array of size n and returns a slice of length 0 and capacity n
	clAgts := make([]restclientagent.RestClientAgent, 0, nAgent)
	servAgt := restserveragent.NewRestServerAgent(url1)

	log.Println("démarrage du serveur...")
	go servAgt.Start()

	log.Println("démarrage des clients...")
	// create 100 agents and append to clAgts table
	for i := 0; i < nAgent; i++ {
		permutation := rand.Perm(nCandidat)
		prefs := make([]rad.Alternative, 10)
		for i := range permutation {
			permutation[i] += 1
			prefs[i] = rad.Alternative(permutation[i])
		}
		id := fmt.Sprintf("id%02d", i)
		agt := restclientagent.NewRestClientAgent(id, url2, prefs)
		clAgts = append(clAgts, *agt)
	}

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		func(agt restclientagent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}

	fmt.Scanln()
	// for i:=10; i>=1; i-- {
	// 	fmt.Printf("Calculation vote in %v", i)
	// 	time.Sleep(1000)
	// }
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"serveur-vote/types"
	"serveur-vote/vote/restclientagent"
	"serveur-vote/vote/restserveragent"
	"time"
)

func main() {
	const nAgent = 100
	const nCandidat = 10
	const url1 = ":8080"
	const url2 = "http://localhost:8080"
	// allocates an underlying array of size n and returns a slice of length 0 and capacity n
	ballotCreationAgt := restclientagent.NewRestClientAgent("ballot_creation_agt", url2, nil, 0)
	clAgts := make([]restclientagent.RestClientAgent, 0, nAgent)
	servAgt := restserveragent.NewRestServerAgent(url1)

	log.Println("démarrage du serveur...")
	go servAgt.Start()
	log.Println("démarrage le client qui génére les ballots...")
	go ballotCreationAgt.CreateBallotStart()

	time.Sleep(5 * time.Second)

	log.Println("démarrage des clients...")
	// create 100 agents and append to clAgts table
	for i := 0; i < nAgent; i++ {
		permutation := rand.Perm(nCandidat)
		prefs := make([]types.Alternative, 10)
		for i := range permutation {
			permutation[i] += 1
			prefs[i] = types.Alternative(permutation[i])
		}
		id := fmt.Sprintf("id%02d", i)
		agt := restclientagent.NewRestClientAgent(id, url2, prefs, rand.Intn(10))
		clAgts = append(clAgts, *agt)
	}

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		func(agt restclientagent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}

	fmt.Scanln()
}

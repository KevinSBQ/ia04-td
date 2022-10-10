package main

import (
	"fmt"
	"log"
	"math/rand"
	"td4/restagentdemo/restclientagent"
	"td4/restagentdemo/restserveragent"
)

func main() {
	const n = 100
	const url1 = ":8080"
	const url2 = "http://localhost:8080"
	ops := [...]string{"+", "-", "*"}
	// allocates an underlying array of size n and returns a slice of length 0 and capacity n
	clAgts := make([]restclientagent.RestClientAgent, 0, n)
	servAgt := restserveragent.NewRestServerAgent(url1)

	log.Println("démarrage du serveur...")
	go servAgt.Start()

	log.Println("démarrage des clients...")
	// create 100 agents and append to clAgts table
	for i := 0; i < n; i++ {
		id := fmt.Sprintf("id%02d", i)
		op := ops[rand.Intn(3)]
		op1 := rand.Intn(100)
		op2 := rand.Intn(100)
		agt := restclientagent.NewRestClientAgent(id, url2, op, op1, op2)
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

package main

import (
	"fmt"

	ras "serveur-vote/vote/restserveragent"
)

func main() {
	server := ras.NewRestServerAgent(":8080")
	server.Start()
	fmt.Scanln()
}

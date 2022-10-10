package main

import (
	"fmt"

	ras "td4/restagentdemo/restserveragent"
)

func main() {
	server := ras.NewRestServerAgent(":8080")
	server.Start()
	fmt.Scanln()
}

package main

import (
	"fmt"

	"td4/restagentdemo/restclientagent"
)

func main() {
	ag := restclientagent.NewRestClientAgent("id1", "http://localhost:8000", "+", 11, 1)
	ag.Start()
	fmt.Scanln()
}

package restclientagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	rad "td4/vote"
)

type RestClientAgent struct {
	id       string
	url      string
	prefs    []rad.Alternative
}

func NewRestClientAgent(id string, url string, prefs []rad.Alternative) *RestClientAgent {
	// constructor of RestClientAgent

	return &RestClientAgent{id, url, prefs}
}

func (rca *RestClientAgent) treatVoteResponse(r *http.Response) string {
	// create a new Response type, transform from json to object, write response

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// Response is a structure which contains only an integer : Result int `json:"res"`
	var resp rad.VoteResponse
	json.Unmarshal(buf.Bytes(), &resp)
	return resp.Result
}

func (rca *RestClientAgent) treatResultResponse(r *http.Response) rad.Alternative {
	// create a new Response type, transform from json to object, write response

	// fmt.Println("httpResponse: ", r)
	// fmt.Println("header: ", r.Header)
	// fmt.Println("body: ", r.Body)

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// Response is a structure which contains only an integer : Result int `json:"res"`
	var resp rad.ResultResponse
	json.Unmarshal(buf.Bytes(), &resp)

	// fmt.Println("Unmarshalled resp: ", resp)

	return resp.Result
}

func (rca *RestClientAgent) treatCountResponse(r *http.Response) int {
	// create a new Response type, transform from json to object, write response

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// Response is a structure which contains only an integer : Result int `json:"res"`
	var resp int
	json.Unmarshal(buf.Bytes(), &resp)

	// fmt.Println("Unmarshalled resp: ", resp)

	return resp
}

func (rca *RestClientAgent) doVoteRequest() (res string, err error) {

	// Request is a structure of an operator string and two int arguments :
	// Operator string `json:"op"`
	// Args [2]int `json:"args"`

	// create a new Request type and allocate it with agent's own operator and arguments
	req := rad.VoteRequest{
		Prefs: rca.prefs,
	}

	// sérialisation de la requête
	url := rca.url + "/vote"
	data, _ := json.Marshal(req)

	// envoi de la requête
	// equivalence dans le diapo : string(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	// res returned from treatResponse, than return it to main
	res = rca.treatVoteResponse(resp)

	return
}

func (rca *RestClientAgent) doResultRequest() (res rad.Alternative, err error) {

	url := rca.url + "/reqresult"

	// envoi de la requête
	// equivalence dans le diapo : string(data)
	resp, err := http.Get(url)
	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	// res returned from treatResponse, than return it to main
	res = rca.treatResultResponse(resp)

	return
}

func (rca *RestClientAgent) doCountRequest() (res int, err error) {

	url := rca.url + "/reqcount"

	// envoi de la requête
	// equivalence dans le diapo : string(data)
	resp, err := http.Get(url)
	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	// res returned from treatResponse, than return it to main
	res = rca.treatCountResponse(resp)

	return
}

func (rca *RestClientAgent) Start() {
	log.Printf("démarrage de %s", rca.id)
	res1, err1 := rca.doVoteRequest()

	if err1 != nil {
		log.Fatal(rca.id, "error:", err1.Error())
	} else {
		log.Printf("[POST][%s] voting, preferences = %v, %s\n", rca.id, rca.prefs, res1)
	}

	time.Sleep(10 * time.Second)
	res2, err2 := rca.doResultRequest()
	if err2 != nil {
		log.Fatal(rca.id, "error:", err2.Error())
	} else {
		log.Printf("[GET][%s] requesting election result, Alternative elected = %v\n", rca.id, res2)
	}

	time.Sleep(15 * time.Second)
	res3, err3 := rca.doCountRequest()
	if err3 != nil {
		log.Fatal(rca.id, "error:", err3.Error())
	} else {
		log.Printf("[GET][%s] requesting server request count, count = %v\n", rca.id, res3)
	}
}

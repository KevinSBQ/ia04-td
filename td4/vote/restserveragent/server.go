package restserveragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"td4/comsoc"
	rad "td4/vote"
)

type RestServerAgent struct {
	sync.Mutex
	id       string
	reqCount int
	profile  rad.Profile
	voteOpen bool
	resultAvailable bool
	result   rad.Alternative
	addr     string
}

func NewRestServerAgent(addr string) *RestServerAgent {
	// Constructor of server agent
	return &RestServerAgent{id: addr, addr: addr, resultAvailable: false, voteOpen: true}
}

// Test de la méthode
func (rsa *RestServerAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	// check request method, if not allowed (not equal to parameter method), send an not allow error using http.ResponseWriter
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

func (*RestServerAgent) decodeRequest(r *http.Request) (req rad.VoteRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// read from request and write the object to req sent sa parameter
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *RestServerAgent) doVote(w http.ResponseWriter, r *http.Request) {
	// create response, calculate the result and writes to response, code the response, writes to writer

	// mise à jour du nombre de requêtes
	// prevent concurrence using mutex package
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++

	if !rsa.voteOpen {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Vote is closed")
		return
	}

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	var resp rad.VoteResponse
	rsa.profile = append(rsa.profile, req.Prefs)

	resp.Result = "vote succesful"
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}

func (rsa *RestServerAgent) doReqcount(w http.ResponseWriter, r *http.Request) {
	// writes the count number to writer
	if !rsa.checkMethod("GET", w, r) {
		return
	}

	w.WriteHeader(http.StatusOK)
	rsa.Lock()
	defer rsa.Unlock()
	serial, _ := json.Marshal(rsa.reqCount)
	w.Write(serial)
}

func (rsa *RestServerAgent) doReqresult(w http.ResponseWriter, r *http.Request) {
	// writes the count number to writer
	if !rsa.checkMethod("GET", w, r) {
		return
	}
	rsa.Lock()
	defer rsa.Unlock()
	if rsa.resultAvailable {
		// make new ResultResponse to return
		// cannot return directly the result value
		var resp rad.ResultResponse
		resp.Result = rsa.result
		w.WriteHeader(http.StatusOK)
		serial, _ := json.Marshal(resp)
		w.Write(serial)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Access denied, no result yet")
	}
}

func (rsa *RestServerAgent) calculateResult(s *http.Server) {

	for i:=5; i>=1; i-- {
		fmt.Printf("vote closing in %v seconds\n", i)
		time.Sleep(1 * time.Second)
	}
	rsa.Lock()
	defer rsa.Unlock()
	rsa.voteOpen = false
	fmt.Printf("calculating result using BORDA SCF ...\n")
	bestAlts, _ := comsoc.BordaSCF(rsa.profile)
	rsa.result, _ = comsoc.TieBreak(bestAlts)
	fmt.Printf("profile = %v\n", rsa.profile)
	fmt.Printf("best alternatives = %v\n", bestAlts)
	fmt.Printf("best alternative after tie break = %v\n", rsa.result)
	rsa.resultAvailable = true
	fmt.Printf("Results are now open for access\n")
}

func (rsa *RestServerAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/vote", rsa.doVote)
	mux.HandleFunc("/reqcount", rsa.doReqcount)
	mux.HandleFunc("/reqresult", rsa.doReqresult)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", rsa.addr)
	go rsa.calculateResult(s)
	go log.Fatal(s.ListenAndServe())
}

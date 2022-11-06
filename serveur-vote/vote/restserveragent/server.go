package restserveragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"serveur-vote/comsoc"
	"serveur-vote/types"
)

type RestServerAgent struct {
	sync.Mutex
	id       string
	reqCount int
	ballots  map[string]types.Ballot
	addr     string
}

func NewRestServerAgent(addr string) *RestServerAgent {
	// Constructor of server agent
	return &RestServerAgent{id: addr, addr: addr}
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

func (*RestServerAgent) decodeRequest(r *http.Request) (req types.VoteRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// read from request and write the object to req sent sa parameter
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*RestServerAgent) decodeResultVoteRequest(r *http.Request) (req types.ResultVoteRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// read from request and write the object to req sent sa parameter
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*RestServerAgent) decodeNewBallotRequest(r *http.Request) (req types.NewBallotRequest, err error) {
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

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprint(w, err.Error())
		return
	}

	bl, ok := rsa.ballots[req.BallotId]
	if !ok { 
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprint(w, "Ballot not found, check ballot id")
		return
	}

	dl, err3 := time.Parse(time.UnixDate, bl.Deadline)
	if err3 != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprint(w, "Attribut deadline not found")
		return
	}
	
	if 	time.Now().After(dl) {
		w.WriteHeader(http.StatusServiceUnavailable) // 503
		fmt.Fprint(w, "Vote is closed")
		return
	}

	// ajout de seuil dans le cas de SCF approval
	if rsa.ballots[req.BallotId].Rule == "approval" {
		bl.Thresholds = append(bl.Thresholds, req.Options)
	}

	// traitement de la requête
	bl.Prof = append(bl.Prof, req.Prefs)
	rsa.ballots[req.BallotId] = bl
	var resp types.VoteResponse
	resp.Result = "vote succesful"
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}

func (rsa *RestServerAgent) doReqNewBallot(w http.ResponseWriter, r *http.Request) {
	// create response, calculate the result and writes to response, code the response, writes to writer

	// mise à jour du nombre de requêtes
	// prevent concurrence using mutex package
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeNewBallotRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	rsa.ballots["vote_"+req.Rule] = types.Ballot{ID: "vote_"+req.Rule, Rule: req.Rule, VoterIds: req.VoterIds, Alts: req.Alts, Deadline: req.Deadline, VoteOpen: true, ResultAvailable: false, Prof: make(types.Profile, 0)}
	// traitement de la requête
	var resp types.VoteResponse
	resp.Result = "ballot created"
	w.WriteHeader(http.StatusCreated)
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
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeResultVoteRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprint(w, err.Error())
		return
	}

	bl, ok := rsa.ballots[req.BallotId]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprint(w, err.Error())
		return
	} else if !bl.ResultAvailable {
		w.WriteHeader(http.StatusTooEarly) // 425
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	rsa.Lock()
	defer rsa.Unlock()
	var resp types.ResultVoteResponse
	resp.Result = bl.Result
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}

func (rsa *RestServerAgent) calculateResultCheck(s *http.Server) {
	rsa.Lock()
	defer rsa.Unlock()
	currentTime := time.Now()
		for i, bl := range rsa.ballots {
			ballotDeadline, err := time.Parse(time.UnixDate, bl.Deadline)
			if err != nil {
				continue
			}
			if currentTime.After(ballotDeadline) && !bl.ResultAvailable {
				log.Printf("[%s] vote closed\n", bl.ID)
				bl.VoteOpen = false
				switch (bl.Rule) {
				case "borda":
					log.Printf("[%s] result calculation started using BORDA SCF\n", bl.ID)
					bestAlts, _ := comsoc.BordaSCF(bl.Prof)
					bl.Result, _ = comsoc.TieBreak(bestAlts)
					// log.Printf("[%s] result calculation, profile = %v\n", bl.ID, bl.Prof)
					log.Printf("[%s] result calculation, best alternatives = %v\n", bl.ID, bestAlts)
					log.Printf("[%s] result calculation, best alternative after tie break = %v\n", bl.ID, bl.Result)
					bl.ResultAvailable = true
					log.Printf("[%s] result calculation completed, best alternative after tie break = %v\n", bl.ID, bl.Result)
					log.Printf("[%s] result open for access", bl.ID)
					rsa.ballots[i] = bl
				case "majority":
					log.Printf("[%s] result calculation started using MAJORITY SCF\n", bl.ID)
					bestAlts, _ := comsoc.MajoritySCF(bl.Prof)
					bl.Result, _ = comsoc.TieBreak(bestAlts)
					// log.Printf("[%s] result calculation, profile = %v\n", bl.ID, bl.Prof)
					log.Printf("[%s] result calculation, best alternatives = %v\n", bl.ID, bestAlts)
					log.Printf("[%s] result calculation, best alternative after tie break = %v\n", bl.ID, bl.Result)
					bl.ResultAvailable = true
					log.Printf("[%s] result calculation completed, best alternative after tie break = %v\n", bl.ID, bl.Result)
					log.Printf("[%s] result open for access", bl.ID)
					rsa.ballots[i] = bl
				case "approval":
					log.Printf("[%s] result calculation started using APPROVAL SCF\n", bl.ID)
					bestAlts, _ := comsoc.ApprovalSCF(bl.Prof, bl.Thresholds)
					bl.Result, _ = comsoc.TieBreak(bestAlts)
					// log.Printf("[%s] result calculation, profile = %v\n", bl.ID, bl.Prof)
					log.Printf("[%s] result calculation, seuils = %v\n", bl.ID, bl.Thresholds)
					log.Printf("[%s] result calculation, best alternatives = %v\n", bl.ID, bestAlts)
					log.Printf("[%s] result calculation, best alternative after tie break = %v\n", bl.ID, bl.Result)
					bl.ResultAvailable = true
					log.Printf("[%s] result calculation completed, best alternative after tie break = %v\n", bl.ID, bl.Result)
					log.Printf("[%s] result open for access", bl.ID)
					rsa.ballots[i] = bl
				case "stv":
					log.Printf("[%s] result calculation started using STV SCF\n", bl.ID)
					bestAlts, _ := comsoc.STV_SCF(bl.Prof)
					bl.Result, _ = comsoc.TieBreak(bestAlts)
					// log.Printf("[%s] result calculation, profile = %v\n", bl.ID, bl.Prof)
					log.Printf("[%s] result calculation, best alternatives = %v\n", bl.ID, bestAlts)
					log.Printf("[%s] result calculation, best alternative after tie break = %v\n", bl.ID, bl.Result)
					bl.ResultAvailable = true
					log.Printf("[%s] result calculation completed, best alternative after tie break = %v\n", bl.ID, bl.Result)
					log.Printf("[%s] result open for access", bl.ID)
					rsa.ballots[i] = bl
				case "kemeny":
					log.Printf("[%s] result calculation started using KEMENY SCF\n", bl.ID)
					bl.Result = comsoc.KemenySCF(bl.Prof)
					// log.Printf("[%s] result calculation, profile = %v\n", bl.ID, bl.Prof)
					log.Printf("[%s] result calculation, best alternative after tie break = %v\n", bl.ID, bl.Result)
					bl.ResultAvailable = true
					log.Printf("[%s] result calculation completed, best alternative after tie break = %v\n", bl.ID, bl.Result)
					log.Printf("[%s] result open for access", bl.ID)
					rsa.ballots[i] = bl
				default:
					break;
				}
				
			}
		}
}


func (rsa *RestServerAgent) calculateResult(s *http.Server) {
	for{
		rsa.calculateResultCheck(s)
		time.Sleep(1 * time.Second)
	}
}

func (rsa *RestServerAgent) Start() {
	rsa.ballots = make(map[string]types.Ballot)
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/vote", rsa.doVote)
	mux.HandleFunc("/reqcount", rsa.doReqcount)
	mux.HandleFunc("/reqresult", rsa.doReqresult)
	mux.HandleFunc("/new_ballot", rsa.doReqNewBallot)

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

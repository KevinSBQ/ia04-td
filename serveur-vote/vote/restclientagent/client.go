package restclientagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"serveur-vote/types"
)

type RestClientAgent struct {
	id       string
	url      string
	prefs    []types.Alternative
	seuil int // valabe uniquement pour ApprovalSCF
}

func NewRestClientAgent(id string, url string, prefs []types.Alternative, seuil int) *RestClientAgent {
	return &RestClientAgent{id, url, prefs, seuil}
}

func (rca *RestClientAgent) treatVoteResponse(r *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var resp types.VoteResponse
	json.Unmarshal(buf.Bytes(), &resp)
	return resp.Result
}

func (rca *RestClientAgent) treatResultVoteResponse(r *http.Response) types.Alternative {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var resp types.ResultVoteResponse
	json.Unmarshal(buf.Bytes(), &resp)
	return resp.Result
}

func (rca *RestClientAgent) treatNewBallotResponse(r *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var resp types.NewBallotResponse
	json.Unmarshal(buf.Bytes(), &resp)
	return resp.BallotId
}

func (rca *RestClientAgent) treatCountResponse(r *http.Response) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var resp int
	json.Unmarshal(buf.Bytes(), &resp)
	return resp
}

func (rca *RestClientAgent) doNewBallotRequest(rule string, deadline string, voterids []string, alts int) (res string, err error) {
	req := types.NewBallotRequest{
		Rule: rule,
		Deadline: deadline,
		VoterIds: voterids,
		Alts: alts,
	}
	url := rca.url + "/new_ballot"
	data, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = rca.treatNewBallotResponse(resp)
	return
}

func (rca *RestClientAgent) doVoteRequest(ballotId string) (res string, err error) {
	req := types.VoteRequest{
		AgentId: rca.id,
		BallotId: ballotId,
		Prefs: rca.prefs,
		Options: rca.seuil,
	}
	url := rca.url + "/vote"
	data, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = rca.treatVoteResponse(resp)
	return
}

func (rca *RestClientAgent) doResultVoteRequest(ballotId string) (res types.Alternative, err error) {
	req := types.ResultVoteRequest{
		BallotId: ballotId,
	}
	url := rca.url + "/reqresult"
	data, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = rca.treatResultVoteResponse(resp)
	return
}

func (rca *RestClientAgent) doCountRequest() (res int, err error) {
	url := rca.url + "/reqcount"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = rca.treatCountResponse(resp)
	return
}

func (rca *RestClientAgent) CreateBallotStart() {
	log.Printf("création de différent ballots ...")
	res1a, err1a := rca.doNewBallotRequest("borda", time.Now().Add(10 * time.Second).Format(time.UnixDate), nil, 100)
	if err1a != nil {
		log.Fatal(rca.id, "error:", err1a.Error())
	} else {
		log.Printf("[POST][%s] request to create ballot, id = %v, %s\n", rca.id, "borda", res1a)
	}
	res1b, err1b := rca.doNewBallotRequest("majority", time.Now().Add(10 * time.Second).Format(time.UnixDate), nil, 100)
	if err1b != nil {
		log.Fatal(rca.id, "error:", err1b.Error())
	} else {
		log.Printf("[POST][%s] request to create ballot, id = %v, %s\n", rca.id, "majority", res1b)
	}
	res1c, err1c := rca.doNewBallotRequest("approval", time.Now().Add(10 * time.Second).Format(time.UnixDate), nil, 100)
	if err1c != nil {
		log.Fatal(rca.id, "error:", err1c.Error())
	} else {
		log.Printf("[POST][%s] request to create ballot, id = %v, %s\n", rca.id, "approval", res1c)
	}
	res1d, err1d := rca.doNewBallotRequest("stv", time.Now().Add(10 * time.Second).Format(time.UnixDate), nil, 100)
	if err1d != nil {
		log.Fatal(rca.id, "error:", err1d.Error())
	} else {
		log.Printf("[POST][%s] request to create ballot, id = %v, %s\n", rca.id, "stv", res1d)
	}
	res1e, err1e := rca.doNewBallotRequest("kemeny", time.Now().Add(10 * time.Second).Format(time.UnixDate), nil, 100)
	if err1e != nil {
		log.Fatal(rca.id, "error:", err1e.Error())
	} else {
		log.Printf("[POST][%s] request to create ballot, id = %v, %s\n", rca.id, "kemeny", res1e)
	}
}

func (rca *RestClientAgent) Start() {
	log.Printf("démarrage de %s", rca.id)
	res1a, err1a := rca.doVoteRequest("vote_borda")
	if err1a != nil {
		log.Fatal(rca.id, "error:", err1a.Error())
	} else {
		log.Printf("[POST][%s] voting to vote_borda, preferences = %v, %s\n", rca.id, rca.prefs, res1a)
	}
	res1b, err1b := rca.doVoteRequest("vote_majority")
	if err1b != nil {
		log.Fatal(rca.id, "error:", err1b.Error())
	} else {
		log.Printf("[POST][%s] voting to vote_majority, preferences = %v, %s\n", rca.id, rca.prefs, res1b)
	}
	res1c, err1c := rca.doVoteRequest("vote_approval")
	if err1c != nil {
		log.Fatal(rca.id, "error:", err1c.Error())
	} else {
		log.Printf("[POST][%s] voting to vote_approval, preferences = %v, %s\n", rca.id, rca.prefs, res1c)
	}
	res1d, err1d := rca.doVoteRequest("vote_stv")
	if err1d != nil {
		log.Fatal(rca.id, "error:", err1d.Error())
	} else {
		log.Printf("[POST][%s] voting to vote_stv, preferences = %v, %s\n", rca.id, rca.prefs, res1d)
	}
	res1e, err1e := rca.doVoteRequest("vote_kemeny")
	if err1e != nil {
		log.Fatal(rca.id, "error:", err1e.Error())
	} else {
		log.Printf("[POST][%s] voting to vote_kemeny, preferences = %v, %s\n", rca.id, rca.prefs, res1e)
	}

	time.Sleep(10 * time.Second)
	
	res2a, err2a := rca.doResultVoteRequest("vote_borda")
	if err2a != nil {
		log.Fatal(rca.id, "error:", err2a.Error())
	} else {
		log.Printf("[GET][%s] requesting result of election [vote_borda], Alternative elected = %v\n", rca.id, res2a)
	}
	res2b, err2b := rca.doResultVoteRequest("vote_majority")
	if err2b != nil {
		log.Fatal(rca.id, "error:", err2b.Error())
	} else {
		log.Printf("[GET][%s] requesting result of election [vote_majority], Alternative elected = %v\n", rca.id, res2b)
	}
	res2c, err2c := rca.doResultVoteRequest("vote_approval")
	if err2c != nil {
		log.Fatal(rca.id, "error:", err2c.Error())
	} else {
		log.Printf("[GET][%s] requesting result of election [vote_approval], Alternative elected = %v\n", rca.id, res2c)
	}
	res2d, err2d := rca.doResultVoteRequest("vote_stv")
	if err2d != nil {
		log.Fatal(rca.id, "error:", err2d.Error())
	} else {
		log.Printf("[GET][%s] requesting result of election [vote_stv], Alternative elected = %v\n", rca.id, res2d)
	}
	res2e, err2e := rca.doResultVoteRequest("vote_kemeny")
	if err2e != nil {
		log.Fatal(rca.id, "error:", err2e.Error())
	} else {
		log.Printf("[GET][%s] requesting result of election [vote_kemeny], Alternative elected = %v\n", rca.id, res2e)
	}

	time.Sleep(10 * time.Second)

	res3, err3 := rca.doCountRequest()
	if err3 != nil {
		log.Fatal(rca.id, "error:", err3.Error())
	} else {
		log.Printf("[GET][%s] requesting server request count, count = %v\n", rca.id, res3)
	}
}

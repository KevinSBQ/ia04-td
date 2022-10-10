package vote

import ()

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

type VoteRequest struct {
	// Operator string `json:"op"`
	// Args     [2]int `json:"args"`
	Prefs []Alternative `json:"prefs"`
}

type VoteResponse struct {
	Result string `json:"voteres"`
}

type ResultResponse struct {
	Result Alternative `json:"resultres"`
}
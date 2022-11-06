package types

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

type Ballot struct {
	ID string
	Prof  Profile
	Rule string
	VoteOpen bool
	ResultAvailable bool
	Result   Alternative
	Deadline string
	VoterIds []string
	Alts int
	Thresholds []int
}

type NewBallotRequest struct {
	Rule string `json:"rule"`
	Deadline string `json:"deadline"`
	VoterIds []string `json:"voterids"`
	Alts int `json:"alts"`
}

type NewBallotResponse struct {
	BallotId string `json:"ballotid"`
}

type VoteRequest struct {
	AgentId string `json:"agentid"`
	BallotId string `json:"ballotid"`
	Prefs []Alternative `json:"prefs"`
	Options int `json:"options"`
}

type VoteResponse struct {
	Result string `json:"voteres"`
}

type ResultVoteRequest struct {
	BallotId string `json:"ballotid"`
}

type ResultVoteResponse struct {
	Result Alternative `json:"resultres"`
}
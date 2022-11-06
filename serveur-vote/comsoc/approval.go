package comsoc

import (
	"serveur-vote/types"
)

func ApprovalSWF(p types.Profile, thresholds []int) (count types.Count, err error) {
	err = checkProfile(p)
	if count == nil {
		count = make(map[types.Alternative]int)
	}	
	for i := range p {		
		for j:=0; j<=thresholds[i]; j++ {
			count[p[i][j]]++
		}
	}
	return
}

func ApprovalSCF(p types.Profile, thresholds []int) (bestAlts []types.Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	bestAlts = maxCount(count)
	return
}

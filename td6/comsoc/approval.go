package comsoc

import (
	. "td6/types"
)

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	err = checkProfile(p)
	if count == nil {
		count = make(map[Alternative]int)
	}	
	for i := range p {		
		for j:=0; j<=thresholds[i]; j++ {
			count[p[i][j]]++
		}
	}
	return
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	bestAlts = maxCount(count)
	return
}

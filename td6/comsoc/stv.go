package comsoc

import (
	. "td6/types"
)

func STV_SWF(p Profile) (Count, error){
	err := checkProfile(p)
	if err != nil {
		return nil, err
	}
	count := make(Count)
	score := make(Count)
	pcopy := make(Profile, len(p))
	copy(pcopy, p)
	for i:=1 ; i<=len(p)-1 ; i++ {
		for _, prefs := range pcopy {
			score[prefs[0]]++
		}
		var elim Alternative
		if len(minCount(score))==1 {
			elim = minCount(score)[0]
		} else {
			elim, _ = TieBreak(minCount(score))
		}
		count[elim]=i
		// remove worst candidate from all over the prefs
		for i := range pcopy {
			pcopy[i] = remove(pcopy[i], elim)
		}
	}
	count[p[0][0]] = len(p)
	return count, nil
}
func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := STV_SWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}
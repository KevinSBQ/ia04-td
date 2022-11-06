package comsoc

import (
	"serveur-vote/types"
)

func STV_SWF(p types.Profile) (types.Count, error){
	err := checkProfile(p)
	if err != nil {
		return nil, err
	}
	count := make(types.Count)
	var score types.Count
	pcopy := make(types.Profile, len(p))
	copy(pcopy, p)
	var elim types.Alternative
	for i:=1 ; i<=len(p[0])-1 ; i++ {
		score = make(types.Count)
		for _, j := range pcopy[0] {
			score[types.Alternative(j)] = 0
		}
		// fmt.Println("score before: ", score)
		for _, prefs := range pcopy {
			score[prefs[0]]++
		}
		// fmt.Println("score after: ", score)
		// fmt.Println("minCount(score): ",minCount(score))
		if len(minCount(score))==1 {
			elim = minCount(score)[0]
		} else {
			elim, _ = TieBreak(minCount(score))
		}
		// fmt.Println("elim: ",elim)
		count[elim] = i
		// fmt.Println("count: ",count)
		// remove worst candidate from all over the prefs
		if len(pcopy[0]) > 1 {
			for i := range pcopy {
				// fmt.Println("remove(pcopy[i], elim): ", remove(pcopy[i], elim))
				pcopy[i] = remove(pcopy[i], elim)
			}
		}
	}
	count[pcopy[0][0]] = len(p)
	return count, nil
}
func STV_SCF(p types.Profile) (bestAlts []types.Alternative, err error) {
	count, err := STV_SWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}
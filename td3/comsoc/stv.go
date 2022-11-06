package comsoc

import ("fmt")

func STV_SWF(p Profile) (Count, error){
	err := checkProfile(p)
	if err != nil {
		return nil, err
	}
	count := make(Count)
	var score Count
	pcopy := make(Profile, len(p))
	copy(pcopy, p)
	for i:=1 ; i<=len(p)-2 ; i++ {
		score = make(Count)
		for j:=1 ; j<=len(p[0]); j++ {
			score[Alternative(j)] = 0
		}
		for _, prefs := range pcopy {
			score[prefs[0]]++
		}
		fmt.Println("score: ", score)
		var elim Alternative
		fmt.Println("minCount(score): ",minCount(score))
		if len(minCount(score))==1 {
			elim = minCount(score)[0]
		} else {
			elim, _ = TieBreak(minCount(score))
		}
		fmt.Println("elim: ",elim)
		count[elim] = i
		fmt.Println("count: ",count)
		// remove worst candidate from all over the prefs
		if len(pcopy[0]) > 1 {
		for i := range pcopy {
			fmt.Println("remove(pcopy[i], elim): ", remove(pcopy[i], elim))
			pcopy[i] = remove(pcopy[i], elim)
		}
		}
	}
	fmt.Println("score: ", score)
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
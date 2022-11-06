package comsoc

import (
	"serveur-vote/types"
)

func MajoritySWF(p types.Profile) (count types.Count, err error) {
	err = checkProfile(p)
	if count == nil {
		count = make(map[types.Alternative]int)
	}
	for _, pref := range p {
		count[pref[0]]++
	}
	return
}

// can be also implemented using maxCount()
func MajoritySCF(p types.Profile) (bestAlts []types.Alternative, err error) {
	count, err := MajoritySWF(p)
	if err != nil {
		return
	}
	bestScore := -1
	// find the maximum score
	for _, score := range count {
		if score > bestScore {
			bestScore = score
		}
	}
	// append all candidates who have the highest score
	for i, score := range count {
		if score == bestScore {
			bestAlts = append(bestAlts, i)
		}
	}

	return
}
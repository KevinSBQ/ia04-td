package comsoc

import (
	"serveur-vote/types"
)

func BordaSWF(p types.Profile) (types.Count, error) {
	err := checkProfile(p)
	count := make(map[types.Alternative]int)	
	for _, prefs := range p {
		highScore := len(p[0]) - 1
		for _, pref := range prefs {
			count[pref] += highScore
			highScore--
		}
	}
	return count, err
}

func BordaSCF(p types.Profile) (bestAlts []types.Alternative, err error) {
	count, err := BordaSWF(p)
	bestAlts = maxCount(count)
	return
}

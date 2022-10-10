package comsoc

import (
	. "td4/vote"
)

func CondorcetWinner(prefs Profile) ([]Alternative, error) {
	err := checkProfile(prefs)
	winner := make([]Alternative, 0)
	for candidat := 1; candidat <= len(prefs[0]); candidat++ {
		isMajPref := true
		for candidatCompare := 1; candidatCompare <= len(prefs[0]); candidatCompare++ {
			if candidat == candidatCompare {
				continue
			}
			candidatScore := 0
			candidatCompareScore := 0
			for _, pref := range prefs {
				if isPref(Alternative(candidat), Alternative(candidatCompare), pref) {
					candidatScore++
				} else {
					candidatCompareScore++
				}
			}
			if candidatScore > candidatCompareScore {
				continue
			} else {
				isMajPref = false
				break
			}
		}
		if isMajPref {
			winner = append(winner, Alternative(candidat))
		}
	}
	return winner, err
}

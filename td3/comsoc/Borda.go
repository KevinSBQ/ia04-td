package comsoc

import (
)

func BordaSWF(p Profile) (Count, error) {
	err := checkProfile(p)
	count := make(map[Alternative]int)	
	for _, prefs := range p {
		highScore := len(p[0]) - 1
		for _, pref := range prefs {
			count[pref] += highScore
			highScore--
		}
	}
	return count, err
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	bestAlts = maxCount(count)
	return
}

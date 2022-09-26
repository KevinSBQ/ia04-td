package comsoc

import (
	"errors"
	"sort"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for i, pref := range prefs {
		if pref == alt {
			return i
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	pref1 := rank(alt1, prefs)
	pref2 := rank(alt2, prefs)
	if pref1 < pref2 {
		return true
	} else {
		return false
	}
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count Count) (bestAlts []Alternative) {
	maxc := 0
	for _, v := range count {
		if v > maxc {
			maxc = v
		}
	}
	for k, v := range count {
		if maxc == v {
			bestAlts = append(bestAlts, k)
		}
	}
	return
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois par préférences
func checkProfile(prefs Profile) error {
	for _, prefsIndividual := range prefs {
		for _, pref := range prefsIndividual {
			if pref == 0 {
				return errors.New("profil non complet")
			}
		}
		prefsIndividualCopy := make([]Alternative, len(prefsIndividual))
		copy(prefsIndividualCopy, prefsIndividual)
		sort.Slice(prefsIndividualCopy, func(m, n int) bool {
			return prefsIndividualCopy[m] < prefsIndividualCopy[n]
		})
		lastPref := Alternative(-1)
		for _, pref := range prefsIndividualCopy {
			if lastPref == pref {
				return errors.New("doublon preference")
			}
		}
	}
	return nil
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	checkProfile(prefs)
	altsCopy := make([]Alternative, len(alts))
	copy(altsCopy, alts)
	sort.Slice(altsCopy, func(m, n int) bool {
		return altsCopy[m] < altsCopy[n]
	})
	lastPref := Alternative(-1)
	for _, pref := range altsCopy {
		if lastPref == pref {
			return errors.New("doublon alts")
		}
	}
	return nil
}
package comsoc

import (
	"errors"
	"sort"
	"serveur-vote/types"
)

// type Alternative int
// type Profile [][]Alternative
// type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt types.Alternative, prefs []types.Alternative) int {
	for i, pref := range prefs {
		if pref == alt {
			return i
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 types.Alternative, prefs []types.Alternative) bool {
	pref1 := rank(alt1, prefs)
	pref2 := rank(alt2, prefs)
	if pref1 < pref2 {
		return true
	} else {
		return false
	}
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count types.Count) (bestAlts []types.Alternative) {
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

func minCount(count types.Count) (worstAlts []types.Alternative) {
	minc := 1000000
	for _, v := range count {
		if v < minc {
			minc = v
		}
	}
	for k, v := range count {
		if minc == v {
			worstAlts = append(worstAlts, k)
		}
	}
	return
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois par préférences
func checkProfile(prefs types.Profile) error {
	for _, prefsIndividual := range prefs {
		for _, pref := range prefsIndividual {
			if pref == 0 {
				return errors.New("profil non complet")
			}
		}
		prefsIndividualCopy := make([]types.Alternative, len(prefsIndividual))
		copy(prefsIndividualCopy, prefsIndividual)
		sort.Slice(prefsIndividualCopy, func(m, n int) bool {
			return prefsIndividualCopy[m] < prefsIndividualCopy[n]
		})
		lastPref := types.Alternative(-1)
		for _, pref := range prefsIndividualCopy {
			if lastPref == pref {
				return errors.New("doublon preference")
			}
			lastPref = pref
		}
	}
	return nil
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs types.Profile, alts []types.Alternative) error {
	checkProfile(prefs)
	altsCopy := make([]types.Alternative, len(alts))
	copy(altsCopy, alts)
	sort.Slice(altsCopy, func(m, n int) bool {
		return altsCopy[m] < altsCopy[n]
	})
	lastPref := types.Alternative(-1)
	for _, pref := range altsCopy {
		if lastPref == pref {
			return errors.New("doublon alts")
		}
		lastPref = pref
	}
	return nil
}
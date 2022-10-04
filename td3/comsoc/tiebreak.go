package comsoc

import "errors"

// en cas d'égalité
// pouvoir utiliser différents critères :
// ex: choisir toujours le premier candidat
// ex: choisir le candidat qui est le plus proche de ...
func TieBreak(alts []Alternative) (alt Alternative, err error) {
	if len(alts) == 0 {
		return -1, errors.New("alternative est vide")
	}
	return alts[0], nil
}

func TieBreakFactory(altsFac []Alternative) func([]Alternative) (Alternative, error) {
	AltPreOrder := make([]Alternative, len(altsFac))
	copy(AltPreOrder, altsFac)
	tiebreak := func(alts []Alternative) (Alternative, error) {
		if len(alts) == 0 {
			return -1, errors.New("alternative est vide")
		}
		result := alts[0]
		for _, a := range alts {
			if rank(a, AltPreOrder) < rank(result, AltPreOrder) {
				result = a
			}
		}
		return result, nil
	}
	return tiebreak
}

func remove(slice []Alternative, s Alternative) []Alternative {
	return append(slice[:s], slice[s+1:]...)
}

// preference global
func SWFFactory(swf func(p Profile) (Count, error), tiebreak func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
	SWFProduct := func(pp Profile) ([]Alternative, error) {
		// use swf method to get each candidates' score
		count, errSWF := swf(pp)
		if errSWF != nil {
			return nil, errSWF
		}
		// create a strict order table to be returned
		orderStrict := make([]Alternative, len(count))
		for {
			bestAlts := maxCount(count)
			if len(bestAlts) == 0 {
				// if no more elements in bestAlts, end of loop
				break
			} else if len(bestAlts) == 1 {
				// if one single alt of high score, add it directly
				// remove it from the count table
				orderStrict = append(orderStrict, bestAlts[0])
				delete(count, bestAlts[0])
			} else {
				// append to order list and remove from bestAlts one after another
				for len(bestAlts) != 0 {
					alt, errTB := tiebreak(bestAlts)
					if errTB != nil {
						return nil, errTB
					}
					orderStrict = append(orderStrict, alt)
					bestAlts = remove(bestAlts, alt)
					// remove all proceeded alts from count map
					delete(count, alt)
				}
			}
		}
		return orderStrict, nil
	}
	return SWFProduct
}

// candidat choisi
func SCFFactory(scf func(p Profile) (Count, error), tiebreak func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	SCFProduct := func(pp Profile) (Alternative, error) {

		count, errSCF := scf(pp)
		bestAlts := maxCount(count)
		bestAlt, errTB := tiebreak(bestAlts)

		if errSCF != nil {
			return -1, errSCF
		}
		if errTB != nil {
			return -1, errTB
		}
		return bestAlt, nil
	}
	return SCFProduct
}

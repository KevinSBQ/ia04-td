package comsoc

import (
	"errors"
	"serveur-vote/types"
)

// en cas d'égalité
// pouvoir utiliser différents critères :
// ex: choisir toujours le premier candidat
// ex: choisir le candidat qui est le plus proche de ...
func TieBreak(alts []types.Alternative) (alt types.Alternative, err error) {
	if len(alts) == 0 {
		return -1, errors.New("alternative est vide")
	}
	return alts[0], nil
}

func TieBreakFactory(altsFac []types.Alternative) func([]types.Alternative) (types.Alternative, error) {
	AltPreOrder := make([]types.Alternative, len(altsFac))
	copy(AltPreOrder, altsFac)
	tiebreak := func(alts []types.Alternative) (types.Alternative, error) {
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

func remove(slice []types.Alternative, s types.Alternative) []types.Alternative {
	var i int
	for i = range slice {
		if slice[i] == s {
			break
		}
	}
	 return append(slice[:i], slice[i+1:]...)
}

// preference global
func SWFFactory(swf func(p types.Profile) (types.Count, error), tiebreak func([]types.Alternative) (types.Alternative, error)) func(types.Profile) ([]types.Alternative, error) {
	SWFProduct := func(pp types.Profile) ([]types.Alternative, error) {
		// use swf method to get each candidates' score
		count, errSWF := swf(pp)
		if errSWF != nil {
			return nil, errSWF
		}
		// create a strict order table to be returned
		orderStrict := make([]types.Alternative, len(count))
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
func SCFFactory(scf func(p types.Profile) (types.Count, error), tiebreak func([]types.Alternative) (types.Alternative, error)) func(types.Profile) (types.Alternative, error) {
	SCFProduct := func(pp types.Profile) (types.Alternative, error) {

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

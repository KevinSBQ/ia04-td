package main

import (
	"fmt"
	"td6/comsoc"
	. "td6/types"
)

type Pair struct{
	e1 int
	e2 int
}

func (p *Pair) Equal(p2 Pair) bool {
	if p.e1 == p2.e1 && p.e2 == p2.e2 {
		return true
	}
	return false
}

func calculDistanceEdition(pref1 []Alternative, pref2 []Alternative) float64 {
	pairs1 := make([]Pair,0)
	pairs2 := make([]Pair,0)
	for i := 0; i < len(pref1); i++ {
		for j := i+1; j<=len(pref1)-1; j++ {
			pairs1 = append(pairs1, Pair{int(pref1[i]), int(pref1[j])})
		}
	}
	for i := 0; i < len(pref1); i++ {
		for j := i+1; j<=len(pref1)-1; j++ {
			pairs2 = append(pairs2, Pair{int(pref2[i]), int(pref2[j])})
		}
	}
	// fmt.Println(pairs1)
	// fmt.Println(pairs2)
	count := 0
	for _, pair1 := range pairs1 {
		for _, pair2 := range pairs2 {
			if pair1.Equal(pair2) {
				count++
				break
			}
		}
	}
	nbdeconcord := len(pairs1) - count
	disEdition := float64((count - nbdeconcord))
	// fmt.Println("distance d'edition = ", disEdition)
	tau := disEdition / float64(len(pairs1))
	// fmt.Println("tau = ", tau)
	return tau
}

func calculDistanceRangeProfile(pref []Alternative, profil Profile) float64 {
	sum := 0.0
	for _, prefprofil := range profil {
		sum += calculDistanceEdition(pref, prefprofil)
	}
	return sum
}

// Kemeny = recherche d'un rangement minimisant la distance entre lui-meme et le profil
// Algorithme = pour chaque rangement de preferences, calculer la distance entre lui et le profil, choisir le profil dont la distance est min

func KemenySWF(profil Profile) []Alternative {
	fmt.Println("Profile: ", profil)
	minDistance := calculDistanceRangeProfile(profil[0], profil)
	minIndex := 0
	distance := 0.0
	for i, prefs := range profil {
		fmt.Println("index", i)
		fmt.Println("Distance de ", prefs, " est ", calculDistanceRangeProfile(prefs, profil))
		if distance = calculDistanceRangeProfile(prefs, profil); distance < minDistance {
			minDistance = distance
			minIndex = i
		}
	}
	return profil[minIndex]
}

func KemenySCF(profil Profile) Alternative {
	return KemenySWF(profil)[0]
}

func tiebreakMajoritySCF(P Profile) Alternative {
	bestAlts, _ := comsoc.MajoritySCF(P)
	if len(bestAlts) != 1 {
		Alt, _ := comsoc.TieBreak(bestAlts)
		return Alt
	} else {
		return bestAlts[0]
	}
}

func tiebreakBordaSCF(P Profile) Alternative {
	bestAlts, _ := comsoc.BordaSCF(P)
	if len(bestAlts) != 1 {
		Alt, _ := comsoc.TieBreak(bestAlts)
		return Alt
	} else {
		return bestAlts[0]
	}
}

func tiebreakCopelandSCF(P Profile) Alternative {
	bestAlts, _ := comsoc.CopelandSCF(P)
	if len(bestAlts) != 1 {
		Alt, _ := comsoc.TieBreak(bestAlts)
		return Alt
	} else {
		return bestAlts[0]
	}
}

// Full permutation generation
// permutation([1, 2, 3]) => [[1, 2, 3] [1, 3, 2] [2, 1, 3] [2, 3, 1] [3, 1, 2] [3, 2, 1]]
func permutations(arr []int)[][]int{
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}

// make an array filled with a range of numbers
// makeRange(1, 5)  =>  [1, 2, 3, 4, 5]
func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func intToAlternative(table [][]int) [][]Alternative {
	newtable := make([][]Alternative, len(table))
	for i:=0; i<len(table); i++ {
		newtable[i] = make([]Alternative, len(table[0]))
	}
	for i := range table {
		for j := range table[i] {
			newtable[i][j] = Alternative(table[i][j])
		}
	}
	return newtable
} 

// renvoyer vainqueur possible d'election et une des completions
func possibleWinners(P Profile) map[Alternative][]Alternative {
	pWs := make(map[Alternative][]Alternative)
	nbCandidat := len(P[0])
	// fmt.Println("nombre de candidats : ", nbCandidat)
	alts := makeRange(1, nbCandidat)
	prefsPossible := intToAlternative(permutations(alts))
	// fmt.Println("prefsPossible : ", prefsPossible)
	for _, prefsChosen := range prefsPossible {
		newP := append(P, prefsChosen)
		winner := tiebreakMajoritySCF(newP)
		if _, ok := pWs[winner]; !ok {
			pWs[winner] = prefsChosen
		}
	}
	return pWs
}

func isPossibleWinner(P Profile, c Alternative) bool {
	pWs := possibleWinners(P)
	if _, ok := pWs[c] ; ok {
		return true
	}
	return false
}

func isNecessaryWinner(P Profile, c Alternative) bool {
	pWs := possibleWinners(P)
	if _, ok := pWs[c] ; ok && len(pWs)==1 {
		return true
	}
	return false
}

func bestResponse(P Profile, pref []Alternative) (winner Alternative, ballot []Alternative) {
	// test de plus préferé à moin préferé
	for _, p := range pref {
		// une fois trouver un candidat possible, quitter le boucle
		if isPossibleWinner(P, p) {
			winner = p
			break
		}
	}
	// on a préservé qu'une seule complétion dans la valeur de retourne de fonction possibleWinners
	// mais il peut y avoir plusieurs
	// il faut revenir à chercher toutes les complétions possibles pour que ce winner soit élu
	possibleCompletion := make([][]Alternative, 0)
	nbCandidat := len(P[0])
	alts := makeRange(1, nbCandidat)
	prefsPossible := intToAlternative(permutations(alts))
	for _, prefsChosen := range prefsPossible {
		newP := append(P, prefsChosen)
		possibleWinner := tiebreakMajoritySCF(newP)
		if possibleWinner == winner {
			possibleCompletion = append(possibleCompletion, prefsChosen)
		}
	}

	fmt.Println("All possible completions : ", possibleCompletion)
	// si une seule possiblité, la renvoyer directement
	if len(possibleCompletion) == 1 {
		ballot = possibleCompletion[0]
		return
	}
	// si plusieurs possiblités, calculer est choisir celui qui a le moins de distance avec pref réel
	minDistance := calculDistanceEdition(possibleCompletion[0], pref)
	minIndex := 0
	for index, completion := range possibleCompletion {
		distance := calculDistanceEdition(completion, pref)
		if distance < minDistance {
			minDistance = distance
			minIndex = index
		}
	}
	ballot = possibleCompletion[minIndex]
	return 
}

func main(){
	alt1 := []Alternative{3,1,2,4}
	alt2 := []Alternative{1,2,3,4}
	fmt.Println(calculDistanceEdition(alt1, alt2))

	profil := [][]Alternative{{3,1,2,4},{3,1,2,4},{1,2,3,4},{3,4,2,1},{2,1,3,4}}
	fmt.Println(calculDistanceRangeProfile(alt1, profil))

	fmt.Println("KemenySWF : ",KemenySWF(profil))
	fmt.Println("KemenySCF : ",KemenySCF(profil))

	profilManipulation := [][]Alternative{{1,2,3,4,5}, {2,1,4,5,3}, {5,3,2,1,4}, {5,3,2,1,4}}
	prefs := []Alternative{4,3,1,2,5}
	fmt.Println("Preference of last voter : ", prefs)
	fmt.Println("Possible Winners : ", possibleWinners(profilManipulation))
	winner, ballot := bestResponse(profilManipulation, prefs)
	fmt.Println("Best Response : ", ballot)
	fmt.Println("Winner in this case : ", winner)
}
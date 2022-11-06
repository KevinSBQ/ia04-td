package comsoc

import (
	"serveur-vote/types"
)

func CopelandSWF(p types.Profile) (types.Count, error){
	err := checkProfile(p)
	if err != nil {
		return nil, err
	}
	count := make(types.Count)
	for i:=1; i<=len(p[0]); i++ {
		for j:=i+1; j<=len(p[0]); i++ {
			prefI := 0
			prefJ := 0
			for _, prefs := range p {
				rankI := rank(types.Alternative(i) ,prefs)
				rankJ := rank(types.Alternative(j) ,prefs)
				if rankI < rankJ {
					prefI++
				} else {
					prefJ++
				}
			}
			if prefI > prefJ {
				count[types.Alternative(i)]++
				count[types.Alternative(j)]--
			} else if prefJ > prefI{
				count[types.Alternative(i)]--
				count[types.Alternative(j)]++
			}
		}
	}
	return count, nil
}

func CopelandSCF(p types.Profile) (bestAlts []types.Alternative, err error){
	count, err := CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count),nil
}
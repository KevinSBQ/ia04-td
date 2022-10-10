package comsoc

import (
	. "td4/vote"
)

func CopelandSWF(p Profile) (Count, error){
	err := checkProfile(p)
	if err != nil {
		return nil, err
	}
	count := make(Count)
	for i:=1; i<=len(p[0]); i++ {
		for j:=i+1; j<=len(p[0]); i++ {
			prefI := 0
			prefJ := 0
			for _, prefs := range p {
				rankI := rank(Alternative(i) ,prefs)
				rankJ := rank(Alternative(j) ,prefs)
				if rankI < rankJ {
					prefI++
				} else {
					prefJ++
				}
			}
			if prefI > prefJ {
				count[Alternative(i)]++
				count[Alternative(j)]--
			} else if prefJ > prefI{
				count[Alternative(i)]--
				count[Alternative(j)]++
			}
		}
	}
	return count, nil
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error){
	count, err := CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count),nil
}
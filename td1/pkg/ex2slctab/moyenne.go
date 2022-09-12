package ex2slctab

import (
)

func Moyenne(sl []int) float64 {
	sum := 0.0
	for i := 0; i < len(sl); i++ {
		sum += float64(sl[i])
	}
	sum /= float64(len(sl))
	return sum
}

func MoyenneRange(sl []int) float64 {
	sum := 0.0
	for _,val := range sl {
		sum += float64(val)
	}
	sum /= float64(len(sl))
	return sum
}

func MoyenneImplicite(sl []int) (moy float64) {
	// moy = 0.0 pas besoin car 0-value of float64 is 0.0
	for _,val := range sl {
		moy += float64(val)
	}
	moy /= float64(len(sl))
	return
}
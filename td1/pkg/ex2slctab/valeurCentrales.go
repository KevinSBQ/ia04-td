package ex2slctab

import(
	"sort"
	"fmt"
)

func ValeurCentrales(sl []int) []int {
	fmt.Printf("\n")
	slorder := make([]int, len(sl))
	copy(slorder, sl)
	sort.Ints(slorder)
	fmt.Println("Apres tri : sl = ", sl, ", slorder = ", slorder)
	if len(slorder) % 2 == 0 {
		return []int{slorder[len(slorder) / 2 - 1], slorder[len(slorder) / 2]}
	} else {
		return []int{slorder[len(slorder) / 2]}
	}
}

func ValeurCentrales2(sl []int) (vc []int) {
	sl2 := make([]int, len(sl)) // make obligatoire pour initialization, sinon, le slice sera nil
	copy(sl2, sl)
	sort.Slice(sl2, func(i, j int) bool {return sl2[i] < sl2[j]})
	if i := len(sl2)/2; len(sl2)%2==0 {
		vc := make([]int, 2)
		copy(vc, sl2[i-1:i+1])
	} else {
		vc := make([]int, 1)
		copy(vc, sl2[i:i+1])
	}
	return
}
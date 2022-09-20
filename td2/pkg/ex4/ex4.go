package ex4

import "sync"

// remplit tab avec la valeur v
func Fill(tab []int, v int){
	for i:= range tab{
		tab[i] = v
	}
}

/*
Repartition de tache :
ex:
tab [11]
wk: 4
step = 11 / 4 = 2
rest = 11 % 4 = 3
works : 		3 3 3 2
worker number : 1 2 3 4
*/

/*

ex2;
tab [14]
wk: 5
step = 14 / 5 = 2
rest = 14 % 5 = 4
works :		3 3 3 3 2
*/

func FillConc(tab []int, v int, workerCount int){
	var(
		tabSize = len(tab)
		step = tabSize / workerCount
		rest = tabSize % workerCount
		firstIndex, lastIndex int
		wg sync.WaitGroup
	)

	wg.Add(workerCount)

	for i:=0 ; i< workerCount; i++ {
		lastIndex += step
		if i < rest {
			lastIndex++
		}
		go func(i, j int){
			Fill(tab[i:j],v)
			wg.Done()
		}(firstIndex, lastIndex)
		firstIndex = lastIndex
	}
	wg.Wait()
}
// applique f sur chaque élément de tab et remplace la valeurÒ
func ForEach(tab []int, f func (int) int){
	for i:= range tab{
		tab[i] = f(tab[i])
	}
}


func ForEachConc(tab []int, f func (int) int, workerCount int){
	var(
		tabSize = len(tab)
		step = tabSize / workerCount
		rest = tabSize % workerCount
		firstIndex, lastIndex int
		wg sync.WaitGroup
	)

	wg.Add(workerCount)

	for i:=0 ; i< workerCount; i++ {
		lastIndex += step
		if i < rest {
			lastIndex++
		}
		go func(i, j int){
			ForEach(tab[i:j], f)
			wg.Done()
		}(firstIndex, lastIndex)
		firstIndex = lastIndex
	}
	wg.Wait()
}

// copy le tableau src dans dest
func Copy(src []int, dest []int){
	for i:= range dest{
		dest[i] = src[i]
	}
}
// // vérifie que tab1 et tab2 sont identiques
func Equal(tab1 []int, tab2 []int) bool {
	if len(tab1) != len(tab2) {
		return false
	}
	for i:=0; i<len(tab1); i++ {
		if tab1[i] != tab2[i] {
			return false
		}
	}
	return true
}

func EqualConc(tab1, tab2 []int, workerCount int) bool {	
	if len(tab1) != len(tab2) {
		return false
	}
	var(
		tabSize = len(tab1)
		step = tabSize / workerCount
		rest = tabSize % workerCount
		firstIndex, lastIndex int
		wg sync.WaitGroup
		workerResults = make([]bool, workerCount)
	)

	wg.Add(workerCount)

	for i:=0 ; i< workerCount; i++ {
		lastIndex += step
		if i < rest {
			lastIndex++
		}
		go func(fi, li, i int){
			workerResults[i] = Equal(tab1[fi:li], tab2[fi:li])
			wg.Done()
		}(firstIndex, lastIndex, i)
		firstIndex = lastIndex
	}
	wg.Wait()

	for i:=0; i<workerCount; i++{
		if !workerResults[i] {
			return false
		}
	}
	return true
}

func EqualConcInterup(tab1, tab2 []int, workerCount int) bool {	
	if len(tab1) != len(tab2) {
		return false
	}
	var(
		tabSize = len(tab1)
		step = tabSize / workerCount
		rest = tabSize % workerCount
		firstIndex, lastIndex int
		done = make(chan struct{})
		completed = make(chan struct{})
		countdown = workerCount
	)

	for i:=0 ; i< workerCount; i++ {
		lastIndex += step
		if i < rest {
			lastIndex++
		}
		go func(fi, li int, c,d chan struct{}){
			// envoyer un n'importe signal dans channel completed a la fin de chaque fontion worker
			defer func() {c <- struct{}{}}()
			for i:= fi; i<li; i++ {
				// check inside each worker at each iteration if there is a signal of inequality sent by other worker
				// if it is the case, terminates the process (interruption)
				// reduing rate reduces the check frequency of process termination criteria
				// it's ok to check the existance of this signal at each iteration, but it cost time.
				if rate := (len(tab1) / 10); rate ==0 || i%rate == 0{
					select {
						case <- d:
							return
						default:
					}
				}
				if tab1[i] != tab2[i] {
					// envoyer un n'importe signal dans channel done pour dire que deux tab n'est pas egale
					d <- struct{}{}
					return
				}
			}
		}(firstIndex, lastIndex, completed, done)
		firstIndex = lastIndex
	}

	for countdown > 0 {
		select {
		case <- done:
			return false
		case <- completed:
			countdown--
		}
	}
	return true
}

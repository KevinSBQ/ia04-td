package probleme

import ()

func IsPalindrome(word string) bool {
	l := len(word)
	for i := 0; i < l/2; i++ {
		if word[i] != word[l-i-1] {
			return false
		}
	}
	return true
}

func Palindromes(words []string) (l []string){
	for _,word := range words {
		if IsPalindrome(word) {
			l = append(l, word)
		}
	}
	return
}
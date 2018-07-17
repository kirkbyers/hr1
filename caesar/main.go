package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create cipher slice
	// Loop input characters
	// pick cipher'd letter
	// reconstruct string
	var length, delta int
	var input, result string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	alphabetStrLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetStrUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLower := []rune(alphabetStrLower)
	alphabetUpper := []rune(alphabetStrUpper)
	cypherAlphabetLower := createCipherAlphabet(alphabetLower, delta)
	cypherAlphabetUpper := createCipherAlphabet(alphabetUpper, delta)

	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetStrLower, ch) >= 0:
			result += string(rotateRune(ch, alphabetLower, cypherAlphabetLower))
		case strings.IndexRune(alphabetStrUpper, ch) >= 0:
			result += string(rotateRune(ch, alphabetUpper, cypherAlphabetUpper))
		default:
			result += string(ch)
		}
	}

	fmt.Println(result)
}

func createCipherAlphabet(inpAlphabet []rune, shift int) []rune {
	if len(inpAlphabet) < 2 {
		panic("Alphabet must be longer that 2 runes")
	}
	outAlphabet := make([]rune, len(inpAlphabet))
	for i := range inpAlphabet {
		toRune := inpAlphabet[(i+shift)%len(inpAlphabet)]
		outAlphabet[i] = toRune
	}
	return outAlphabet
}

func rotateRune(s rune, inApha []rune, outAlpha []rune) rune {
	idx := strings.IndexRune(string(inApha), s)
	if idx < 0 {
		return s
	}

	return outAlpha[idx]
}

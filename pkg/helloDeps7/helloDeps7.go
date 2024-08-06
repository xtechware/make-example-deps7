package helloDeps7

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps7.PrintPhrase")
	return phrase
}

package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	words := strings.Split(s, " ")
	itog := ""
	for i := len(words) - 1; i >= 0; i-- {
		itog += words[i] + " "
	}
	return strings.TrimSpace(itog)
}

func main() {
	fmt.Println(reverse("snow dog sun"))
}
package main

import "fmt"

func square(a int) int { 
	return a * a
}

func main() {
	chisla := []int{2, 4, 6, 8, 10} 
	chanel := make(chan int)

	for _, chislo := range chisla { 
		go func(a int) { 
			chanel <- square(a) 
		}(chislo) 
	}

	for i := 0; i < len(chisla); i++ { 
		fmt.Println(<-chanel)
	}
}
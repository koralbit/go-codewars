package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Primes in numbers
https://www.codewars.com/kata/54d512e62a5e54c96200019e/train/go
*/

func main()  {
	fmt.Println(PrimeFactors(86240))
	fmt.Println(PrimeFactors(7775460))
	fmt.Println(PrimeFactors(79340))
}

func PrimeFactors(n int) string {
	p := make(map[int]int)
    for i := 2; i <= n; i++ {
		for n % i == 0 {
			n /= i
			p[i]++
		}
	}
	
	keys := make([]int,0)
    for k := range p {
        keys = append(keys, k)
    }
    sort.Ints(keys)
	builder := strings.Builder{}
	for _,k := range keys {
		if p[k] == 1 {
			builder.WriteString(fmt.Sprintf("(%d)", k))
		} else {
			builder.WriteString(fmt.Sprintf("(%d**%d)", k, p[k]))
		}
	}
	return builder.String()
}
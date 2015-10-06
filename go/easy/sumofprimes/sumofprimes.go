package main

import (
	"fmt"
)

func main() {
	sum := 0
	prime_counter := 0
	iterator := 2
	for prime_counter < 1000 {
		tmp, isprime := IsPrime(iterator)
		if isprime {
			prime_counter++
			sum += tmp
//			fmt.Println(tmp)
		}
		iterator++
	}
	fmt.Println(sum)
}

func IsPrime(val int) (toadd int, ret bool) {
	ret = true
	toadd = val
	if val <= 1 {
		ret = false
		toadd = 0
		return
	}
	if val%2 == 0 {
		for i := 2; i <= val/2; i++ {
			if val%i == 0 {
				ret = false
				toadd = 0
				return
			}
		}
	} else {
		for i := 2; i <= (val+1)/2; i++ {
			if val%i == 0 {
				ret = false
				toadd = 0
				return
			}
		}
	}
	return
}

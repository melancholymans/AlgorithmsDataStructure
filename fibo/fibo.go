package main

import (
	"fmt"
	"time"
)

var memo [50]int

func main(){
	N := 49
	for i:=0;i<50;i++ {
		memo[i] = -1
	} 
    now := time.Now()
	fmt.Println(fibo(N))
    fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
	fmt.Println(memo[49])
}

func fibo(N int) int{
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	if memo[N] != -1 {
		return memo[N]
	}
	memo[N] = fibo(N-1)+fibo(N-2)
	return memo[N]
	//return fibo(N-1)+fibo(N-2)
}


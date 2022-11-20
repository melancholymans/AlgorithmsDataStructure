package main

import(
	"fmt"
	"math"
)

var INF = 1000.0

func main(){
	N := 7
	h := []float64{2,9,4,5,1,6,10}
	var dp [7]float64
	for i:=0;i<N;i++ {
		dp[i] = INF
	}
	dp[0] = 0.0
	for i:=1;i<N;i++ {
		if i == 1 {
			dp[i] = math.Abs(h[i]-h[i-1])
		} else {
			dp[i] = math.Min(dp[i-1]+math.Abs(h[i]-h[i-1]),
							 dp[i-2]+math.Abs(h[i]-h[i-2]))
		}
	}
	fmt.Println(dp[N-1])
}
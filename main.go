package main

import "fmt"

// func main() {
// 	primeCount := 0
// 	oddPrimeCount := 0

// 	for i := 1; i <= 6; i++ {
// 		if isPrime(i) {
// 			primeCount++
// 			if i%2 == 1 {
// 				oddPrimeCount++
// 			}
// 		}
// 	}

// 	probability := float64(oddPrimeCount) / float64(primeCount)
// 	fmt.Printf("한 개의 주사위를 던졌을 때, 나온 눈이 소수이면서 홀수일 확률은 %.2f입니다.\n", probability)
// }

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	pPositiveGivenRealPositive := 0.9
	pPositiveGivenRealNegative := 0.2
	pRealPositive := 0.6

	pPositive := pPositiveGivenRealPositive*pRealPositive + pPositiveGivenRealNegative*(1.0-pRealPositive)
	pRealPositiveGivenPositive := (pPositiveGivenRealPositive * pRealPositive) / pPositive
	fmt.Printf("이 과일이 실제로 상품일 때, 상품으로 판정될 확률은 %.2f입니다.\n", pRealPositiveGivenPositive)
}

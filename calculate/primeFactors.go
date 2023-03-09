package calculate

// func PrimeFactors(n int) []int {
// 	factors := make([]int, 0)
// 	for n%2 == 0 {
// 		factors = append(factors, 2)
// 		n = n / 2
// 	}
// 	for i := 3; i*i <= n; i = i + 2 {
// 		for n%i == 0 {
// 			factors = append(factors, i)
// 			n = n / i
// 		}
// 	}
// 	if n > 2 {
// 		factors = append(factors, n)
// 	}
// 	return factors
// }

// PrimeFactors get all prime factors of a number
func PrimeFactors(n int) []int {
	factors := map[int]bool{}
	for n%2 == 0 {
		factors[2] = true
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			factors[i] = true
			n = n / i
		}
	}
	if n > 2 {
		factors[n] = true
	}
	ret := []int{}
	for k := range factors {
		ret = append(ret, k)
	}
	return ret
}

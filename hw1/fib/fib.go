package fib

func FibonacciGen() func() int {
	prev1, prev2, c := 0, 1, 0

	return func() int {
		c++
		if (c > 2) {
			t := prev1
			prev1 = prev2
			prev2 += t
			return prev2
		} else if (c == 1) {
			return 0
		} 
		
		return 1
	}
}
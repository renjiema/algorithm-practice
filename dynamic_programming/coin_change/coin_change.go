package coin_change

var coins []int

func ViolenceCoinChange(cs []int, amount int) int {
	coins = cs
	return violenceDp(amount)
}

func violenceDp(amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	// 求最小值，初始化为amount+1
	res := amount + 1
	for _, coin := range coins {
		subproblem := violenceDp(amount - coin)
		if subproblem == -1 {
			continue
		}
		res = min(res, subproblem+1)

	}
	if res == amount+1 {
		return -1
	}
	return res
}

func min(vals ...int) int {
	var min int
	for _, val := range vals {
		if min == 0 || val <= min {
			min = val
		}
	}
	return min
}

func CacheCoinChange(cs []int, amount int) int {
	coins = cs
	cache := make(map[int]int)
	return cacheDp(cache, amount)
}

func cacheDp(cache map[int]int, amount int) int {
	// 查找缓存
	res, ok := cache[amount]
	if ok {
		return res
	}
	// base case
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	res = amount + 1
	for _, coin := range coins {
		subres := cacheDp(cache, amount-coin)
		if subres == -1 {
			continue
		}
		res = min(res, subres+1)
	}
	if res == amount+1 {
		res = -1
	}
	cache[amount] = res
	return res
}

func ToTopCoinChange(cs []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var lens = amount + 1
	var dp = make([]int, lens)
	// base case
	dp[0] = 0
	for i := 1; i < lens; i++ {
		dp[i] = lens
		for _, coin := range cs {
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == lens {
		return -1
	}

	return dp[amount]
}

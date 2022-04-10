package tradestocks

// 买卖股票：含冷冻期
// link:https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		// dp[i][0]:今天结束后，手中无股票时最大的收益 =>
		// 昨天是冷冻期
		// 昨天没有进行买入，沿用昨天的状态

		dp[i][0] = max(dp[i-1][0], dp[i-1][2])

		// dp[i][1]:今天结束后，手中有股票 =>
		// 昨天买入了
		// 昨天没动静，沿用昨天有股票时的状态
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])

		// dp[i][2]:冷冻期（今天结束后，也就是明天是不能买卖股票的）=>
		// 今天卖出
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

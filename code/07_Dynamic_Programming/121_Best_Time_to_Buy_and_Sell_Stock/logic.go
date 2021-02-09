package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		currentPrice := prices[i]
		if minPrice > currentPrice {
			// 最小価格を設定
			minPrice = currentPrice
			continue
		}

		// 利益を計算し、最大利益を更新
		profit := currentPrice - minPrice
		if maxProfit < profit {
			maxProfit = profit
		}
	}
	return maxProfit
}

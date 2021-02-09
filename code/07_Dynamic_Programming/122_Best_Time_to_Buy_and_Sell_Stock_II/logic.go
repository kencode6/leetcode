package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	buyPrice := prices[0]
	profitSum := 0
	for i := 0; i < len(prices); i++ {
		if buyPrice > prices[i] {
			// より安い価格があれば買取価格を変更
			buyPrice = prices[i]
			continue
		}

		if i < len(prices)-1 {
			// 次の日の価格が当日を下回ったら売り時
			if prices[i] > prices[i+1] {
				// 利益を加算してbuyPriceを次の日の価格に設定するためのフラグを立てる。
				profit := prices[i] - buyPrice
				profitSum += profit

				// 次の日の価格を設定
				buyPrice = prices[i+1]
			}
		} else {
			// 最終日
			if prices[i] > buyPrice {
				profit := prices[i] - buyPrice
				profitSum += profit
			}
		}
	}
	return profitSum
}

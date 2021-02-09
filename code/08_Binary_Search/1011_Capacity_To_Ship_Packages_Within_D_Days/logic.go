package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/

func shipWithinDays(weights []int, d int) int {
	maxWeight := 0 // 全てのweightの合算
	minWeight := 0 // weights内の最大weight
	for _, weight := range weights {
		maxWeight += weight
		if minWeight < weight {
			minWeight = weight
		}
	}

	retWeight := 0
	count := 1
	for {
		if minWeight > maxWeight {
			break
		}
		// 中央重量を計算(二分探索は通常はindexの挟み込みだが概念を拡張すると数量にも利用できる)
		midWeight := (minWeight + maxWeight) / 2
		count++

		// midWeightで分割した場合の必要日数を計算
		requireDays, requireMaxMidWeight := calcRequireDays(weights, midWeight)

		if requireDays > d {
			// 必要日数が上限日数より多い場合、最小重量を中央重量+1に
			minWeight = midWeight + 1
		} else {
			// 必要日数が上限日数より少ない場合、最大重量を中央重量-1にして計算した重量を答えの候補として登録
			maxWeight = midWeight - 1
			retWeight = requireMaxMidWeight

			// ※midWeightはrequireMaxWeightに収束するのでmidWeightを返却しても良い
			// retWeight = midWeight
		}
	}
	return retWeight
}

// midWeightで分割した際の必要日数と最大重量を返却する
func calcRequireDays(weights []int, midWeight int) (int, int) {
	maxWeight := 0
	requireDays := 1
	currentWeight := 0
	for _, weight := range weights {
		currentWeight += weight
		if currentWeight > midWeight {
			// 重量上限を超えたら最大重量を更新する。加算する一つ前の重量が最大重量となる。
			targetWeight := currentWeight - weight
			if maxWeight < targetWeight {
				maxWeight = targetWeight
			}

			// 必要日数を更新し、currentWeightをこの重量から再度計算
			requireDays++
			currentWeight = weight
		}
	}

	if maxWeight < currentWeight {
		maxWeight = currentWeight
	}
	return requireDays, maxWeight
}

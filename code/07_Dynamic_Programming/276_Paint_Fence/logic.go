package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/paint-fence/

func numWays(n int, k int) int {
	// 同じ色は2回までなので、以下の漸化式を満たすS(n)を返却する
	// n=1 S(1) = k
	// n=2 S(2) = k^2
	// n>=3 S(n) = S(n-1) * (k-1) + S(n-2) * (k-1)

	sums := []int{}
	sums = append(sums, k)   // n = 1
	sums = append(sums, k*k) // n = 2

	if n == 0 {
		return 0
	} else if n == 1 {
		return sums[0]
	} else if n == 2 {
		return sums[1]
	}

	for i := 2; i < n; i++ {
		sum := (sums[i-1] + sums[i-2]) * (k - 1)
		sums = append(sums, sum)
	}
	return sums[n-1]
}

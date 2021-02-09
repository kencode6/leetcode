package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/k-th-symbol-in-grammar/

/*
以下のような木構造を検討する
                        0
            ┌───────────┴───────────┐
            0                       1
      ┌─────┴─────┐           ┌─────┴─────┐
      0           1           1           0
   ┌──┴──┐     ┌──┴──┐     ┌──┴──┐     ┌──┴──┐
   0     1     1     0     1     0     0     1

k番目のbit符号は前のbit符号から取得できるので再帰的に計算することで求められる。
*/
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	prevK := (k-1)/2 + 1
	position := k % 2 // left:1, right:0
	bit := kthGrammar(n-1, prevK)

	// 前の値が0の場合
	if bit == 0 {
		if position == 1 {
			return 0
		}
		return 1
	}

	// 1の場合
	if position == 1 {
		return 1
	}
	return 0

	// 上記はこの一行でも満たせる。
	// return (position + bit) % 2
}

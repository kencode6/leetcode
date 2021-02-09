package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 && len(nums2) == 0 {
		return [][]int{}
	}

	pairsContainer := NewPairsContainer(k)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			pair := []int{nums1[i], nums2[j]}
			pairsContainer.Add(pair)
		}
	}
	return pairsContainer.Pairs()
}

// PairsContainer pairの合計が小さい順に要素数limitNumまでのpair要素を格納します。
type PairsContainer struct {
	pairs    [][]int
	limitNum int
}

// NewPairsContainer PairsContainerを生成します。
func NewPairsContainer(limitNum int) *PairsContainer {
	return &PairsContainer{limitNum: limitNum}
}

// Pairs pairsを返します
func (p *PairsContainer) Pairs() [][]int {
	return p.pairs
}

// Add pairを追加します。
func (p *PairsContainer) Add(newPair []int) {
	if len(p.pairs) == 0 {
		//　初回
		p.pairs = append(p.pairs, newPair)
		return
	}

	// add要素と既存要素の比較
	newSum := calcPairSum(newPair)

	// 要素挿入位置を探索
	insertIndex := -1
	for i := 0; i < len(p.pairs); i++ {
		sum := calcPairSum(p.pairs[i])
		if newSum <= sum {
			insertIndex = i
			break
		}
	}

	if insertIndex == -1 {
		// 挿入位置無し
		if len(p.pairs) < p.limitNum {
			p.pairs = append(p.pairs, newPair)
		}
		return
	}

	// 要素の挿入
	p.pairs = append(p.pairs, []int{})
	copy(p.pairs[insertIndex+1:], p.pairs[insertIndex:])
	p.pairs[insertIndex] = newPair

	if len(p.pairs) > p.limitNum {
		// 要素が上限を超えた場合は取り除く
		p.pairs = p.pairs[:p.limitNum]
	}
}

func calcPairSum(pair []int) int {
	return pair[0] + pair[1]
}

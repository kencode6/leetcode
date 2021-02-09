package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/kth-largest-element-in-a-stream/submissions/

type KthLargest struct {
	kth  int
	nums []int
}

func newKthLargest(kth int) *KthLargest {
	return &KthLargest{
		kth:  kth,
		nums: []int{},
	}
}

// Constructor KthLargestを作成します。
// ※Goではこのような命名は一般的ではないが、leetcodeで指定された名前の為
func Constructor(k int, nums []int) KthLargest {
	kthLargest := newKthLargest(k)
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return *kthLargest
}

func (k *KthLargest) Add(val int) int {

	if len(k.nums) == 0 {
		k.nums = append(k.nums, val)
		return k.kVal()
	}

	if len(k.nums) == 1 {
		if val < k.nums[0] {
			k.nums = []int{val, k.nums[0]}
		} else {
			k.nums = append(k.nums, val)
		}
		return k.kVal()
	}

	insertIndex := -1
	for i := 0; i < len(k.nums); i++ {
		sNum := 0
		if i == 0 {
			sNum = -10001
		} else {
			sNum = k.nums[i-1]
		}
		eNum := k.nums[i]
		if sNum <= val && val <= eNum {
			insertIndex = i
			break
		}
	}

	if insertIndex > -1 {
		k.nums = append(k.nums, 0)
		copy(k.nums[insertIndex+1:], k.nums[insertIndex:])
		k.nums[insertIndex] = val
	} else {
		k.nums = append(k.nums, val)
	}
	return k.kVal()
}

func (k *KthLargest) kVal() int {
	if k.kth <= len(k.nums) {
		return k.nums[len(k.nums)-k.kth]
	}
	return -1
}

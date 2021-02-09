package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/linked"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// listnodeをsliceに変換
	nums1 := convListNodeToNums(l1)
	nums2 := convListNodeToNums(l2)

	// 加算処理を実行
	nums := addNums(nums1, nums2)

	// sliceをlistnodeに変換
	retLn := convNumsToListNode(nums)
	return retLn
}

// NewListNode ListNodeを生成します。
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

// convListNodeToNums listnodeをsliceに変換します。
func convListNodeToNums(ln *ListNode) []int {
	nums := []int{}
	for {
		nums = append(nums, ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}
	fmt.Printf("nums:%v\n", nums)
	return nums
}

// addNums sliceの加算処理を行います。
func addNums(nums1 []int, nums2 []int) []int {
	numsLen := len(nums1)
	if numsLen < len(nums2) {
		numsLen = len(nums2)
	}

	sumNums := []int{}
	plus := 0
	for i := 0; i < numsLen; i++ {
		num1 := getVal(nums1, i)
		num2 := getVal(nums2, i)
		num := num1 + num2 + plus
		if num >= 10 {
			num -= 10
			plus = 1
		} else {
			plus = 0
		}
		sumNums = append(sumNums, num)
	}
	if plus == 1 {
		sumNums = append(sumNums, 1)
	}
	fmt.Printf("sumNums:%v\n", sumNums)
	return sumNums
}

func getVal(nums []int, index int) int {
	if index < len(nums) {
		return nums[index]
	}
	return 0
}

// convNumsToListNode sliceをlistnodeに変換します。
func convNumsToListNode(nums []int) *ListNode {
	var headLn *ListNode
	var ln *ListNode
	for _, num := range nums {
		newLn := NewListNode(num)
		if headLn == nil {
			headLn = newLn
		} else {
			ln.Next = newLn
		}
		ln = newLn
	}
	return headLn
}

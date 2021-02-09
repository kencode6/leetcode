package linked

// leetcodeサーバーでなくローカル環境でコードを実行するためのListNode処理用ツールを提供します。

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode leetcode用ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// SliceToListNode numsからListNodeを作成します。
func SliceToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var rootLn *ListNode
	var prevLn *ListNode
	for _, num := range nums {

		ln := newListNode(num)

		if prevLn != nil {
			prevLn.Next = ln
		} else {
			rootLn = ln
		}
		prevLn = ln
	}
	return rootLn
}

// SliceToCycleListNode numsからListNodeを作成し、指定したindexでCycleするListNodeを作成します。
// cycleLn: サイクルしたListNode
// lastLn: 終端のListNode
// jointLn: 接続部分のListNode
func SliceToCycleListNode(nums []int, cycleIndex int) (cycleLn *ListNode, lastLn *ListNode, jointLn *ListNode) {
	if len(nums) == 0 {
		return nil, nil, nil
	}

	var rootLn *ListNode
	var prevLn *ListNode
	for i, num := range nums {
		ln := newListNode(num)

		if prevLn != nil {
			prevLn.Next = ln
		} else {
			rootLn = ln
		}
		prevLn = ln

		if i == cycleIndex {
			jointLn = ln
		}
	}

	// 終了時に終端のNextをcycle対象のListNodeに接続
	prevLn.Next = jointLn

	cycleLn = rootLn
	lastLn = prevLn
	return
}

// ToSlice ListNodeをsliceに変換します。
func (l *ListNode) ToSlice() []int {
	ln := l
	nums := []int{}

	if l == nil {
		return nums
	}

	for {
		nums = append(nums, ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}
	return nums
}

// ToString ListNodeをstringに変換します
func (l *ListNode) ToString() string {
	nums := l.ToSlice()

	var sb strings.Builder
	fmt.Fprint(&sb, "[")
	for i, num := range nums {
		if i > 0 {
			fmt.Fprint(&sb, ", ")
		}
		fmt.Fprint(&sb, strconv.Itoa(num))
	}
	fmt.Fprint(&sb, "]")

	return sb.String()
}

// AtIndex targetIndexと一致するListNodeを返却します。
func (l *ListNode) AtIndex(targetIndex int) *ListNode {
	if targetIndex < 0 {
		return nil
	}
	ln := l
	index := 0
	for ln != nil {
		if targetIndex == index {
			return ln
		}

		ln = ln.Next
		index++
	}
	return nil
}

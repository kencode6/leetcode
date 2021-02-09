package tree

import (
	"fmt"
	"math"
	"strconv"
)

// TreeNode leetcode用2分木ノード
// ※leetcode問題では TreeNodeの構造体のみ使用可能
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// newTreeNode leetcode用二分木ノードを生成
func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// MaxDepth TreeNodeの最大深さを返却します
func (t *TreeNode) MaxDepth() int {
	if t == nil {
		return 0
	}
	leftDepth := t.Left.MaxDepth()
	rightDepth := t.Right.MaxDepth()
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

/*
ToPointerNums tree構造をint pointerに変換します。
例)
入力値:
               3
   ┌───────────┴───────────┐
   9                      20
                     ┌─────┴─────┐
                    15           7
                  ┌──┴──┐
                 11    12

出力値:
[3,9,20,null,null,15,7]
*/
func (t *TreeNode) ToPointerNums() []*int {
	if t == nil {
		return []*int{}
	}

	pNums := []*int{&t.Val}

	// 処理対象のnodeをスタックに入れる
	queue := newTreeNodeDataQueue()
	queue.add(t, 1)

	for {
		currentTnd := queue.pop()
		if currentTnd == nil {
			// queueデータが無かったら終了
			break
		}
		fmt.Printf("Val:%d, depth:%d\n", currentTnd.treeNode.Val, currentTnd.depth)
		currentTn := currentTnd.treeNode

		// 左右を登録
		pNums = append(pNums, getPVal(currentTn.Left))
		pNums = append(pNums, getPVal(currentTn.Right))

		if currentTn.Left != nil {
			// 左を探索queueに追加
			queue.add(currentTn.Left, currentTnd.depth+1)
		}

		if currentTn.Right != nil {
			// 右を探索queueに追加
			queue.add(currentTn.Right, currentTnd.depth+1)
		}
	}

	// 後方のnullを削除
	lastIndex := len(pNums) - 1
	for i := len(pNums) - 1; i >= 0; i-- {
		pNum := pNums[i]
		if pNum != nil {
			lastIndex = i
			break
		}
	}
	pNums = pNums[:lastIndex+1]
	return pNums
}

func getPVal(tn *TreeNode) *int {
	if tn == nil {
		return nil
	}
	return &tn.Val
}

// PrintTree ノードを標準出力します
func (t *TreeNode) PrintTree() {
	lines := [][]string{}
	level := []*TreeNode{}
	next := []*TreeNode{}

	level = append(level, t)

	nn := 1

	widest := 0

	for nn != 0 {
		line := []string{}

		nn = 0

		for _, n := range level {
			if n == nil {
				line = append(line, "")

				next = append(next, nil)
				next = append(next, nil)
			} else {
				aa := strconv.Itoa(n.Val)
				line = append(line, aa)
				if len(aa) > widest {
					widest = len(aa)
				}

				next = append(next, n.Left)
				next = append(next, n.Right)

				if n.Left != nil {
					nn++
				}
				if n.Right != nil {
					nn++
				}
			}
		}

		if widest%2 == 1 {
			widest++
		}
		lines = append(lines, line)

		tmp := level
		level = next
		next = tmp
		next = []*TreeNode{}
	}

	perpiece := len(lines[len(lines)-1]) * (widest + 4)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		hpw := int(math.Floor(float64(perpiece)/2.0) - 1)

		if i > 0 {
			for j := 0; j < len(line); j++ {

				// split node
				c := " "
				if j%2 == 1 {
					if line[j-1] != "" {
						if line[j] != "" {
							c = "┴"
						} else {
							c = "┘"
						}
					} else {
						if j < len(line) && line[j] != "" {
							c = "└"
						}
					}
				}
				fmt.Print(c)

				// lines and spaces
				if line[j] == "" {
					for k := 0; k < perpiece-1; k++ {
						fmt.Print(" ")
					}
				} else {
					for k := 0; k < hpw; k++ {
						if j%2 == 0 {
							fmt.Print(" ")
						} else {
							fmt.Print("─")
						}
					}
					if j%2 == 0 {
						fmt.Print("┌")
					} else {
						fmt.Print("┐")
					}
					for k := 0; k < hpw; k++ {
						if j%2 == 0 {
							fmt.Print("─")
						} else {
							fmt.Print(" ")
						}
					}
				}
			}
			fmt.Println()
		}

		// print line of numbers
		for j := 0; j < len(line); j++ {

			f := line[j]
			gap1 := int(math.Ceil(float64(perpiece)/2.0 - float64(len(f))/2.0))
			gap2 := int(math.Floor(float64(perpiece)/2.0 - float64(len(f))/2.0))

			// a number
			for k := 0; k < gap1; k++ {
				fmt.Print(" ")
			}
			fmt.Print(f)
			for k := 0; k < gap2; k++ {
				fmt.Print(" ")
			}
		}
		fmt.Println()

		perpiece /= 2
	}
}

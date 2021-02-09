package tree

import (
	"fmt"
)

// treeBuilder TreeNode生成用のビルダー
type treeBuilder struct {
}

// newTreeBuilder 要素を指定してTreeBuilderを生成する。
func newTreeBuilder() *treeBuilder {
	return &treeBuilder{}
}

// treeNodeData 生成過程のTreeNode情報保持用
type treeNodeData struct {
	treeNode *TreeNode
	depth    int
}

func newTreeNodeData(tn *TreeNode, depth int) *treeNodeData {
	return &treeNodeData{
		treeNode: tn,
		depth:    depth,
	}
}

// treeNodeDataQueue 生成過程のTreeNode情報保持用キュー
type treeNodeDataQueue struct {
	treeNodeDatas []*treeNodeData
}

func newTreeNodeDataQueue() *treeNodeDataQueue {
	return &treeNodeDataQueue{
		treeNodeDatas: []*treeNodeData{},
	}
}

func (q *treeNodeDataQueue) add(tn *TreeNode, depth int) {
	tnd := newTreeNodeData(tn, depth)
	q.treeNodeDatas = append(q.treeNodeDatas, tnd)
}

func (q *treeNodeDataQueue) pop() *treeNodeData {
	if len(q.treeNodeDatas) == 0 {
		return nil
	}
	tnd := q.treeNodeDatas[0]
	q.treeNodeDatas = q.treeNodeDatas[1:]
	return tnd
}

/*
CreateTreeNode ポインタ型int sliceから TreeNodeを作成します。
例)
入力値:
[3,9,20,null,null,15,7]

出力値:
               3
   ┌───────────┴───────────┐
   9                      20
                     ┌─────┴─────┐
                    15           7
                  ┌──┴──┐
                 11    12
*/
func CreateTreeNode(pNums []*int) *TreeNode {
	tn, err := newTreeBuilder().Build(pNums)
	if err != nil {
		panic(err)
	}
	return tn
}

// Build ポインタ型int sliceから TreeNodeを作成します。
func (t *treeBuilder) Build(pNums []*int) (*TreeNode, error) {
	if len(pNums) == 0 {
		return nil, nil
	} else if pNums[0] == nil {
		return nil, fmt.Errorf("要素の初めはnull要素にできません。")
	}

	// 処理対象のnodeをスタックに入れる
	root := newTreeNode(*pNums[0])
	queue := newTreeNodeDataQueue()
	queue.add(root, 1)

	index := 0

	for {
		currentTnd := queue.pop()
		if currentTnd == nil {
			return nil, fmt.Errorf("未処理のデータがあります。引数のnumsを見直して下さい。 生成中のTreeNode:%v 未処理のデータ:%v", root, pNums[index:])
		}
		// fmt.Printf("Val:%d, depth:%d\n", currentTnd.treeNode.Val, currentTnd.depth)
		currentTn := currentTnd.treeNode

		index++
		pNum, ok := getPNum(pNums, index)
		if !ok {
			// 要素を全てTreeNodeに登録し終わったので完了
			break
		}
		if pNum != nil {
			// 左のTreeNodeを作成し、次の処理対象としてqueueに追加
			newTn := newTreeNode(*pNum)
			currentTn.Left = newTn
			queue.add(newTn, currentTnd.depth+1)
		}

		index++
		pNum, ok = getPNum(pNums, index)
		if !ok {
			// 要素を全てTreeNodeに登録し終わったので完了
			break
		}
		if pNum != nil {
			// 右のTreeNodeを作成し、次の処理対象としてqueueに追加
			newTn := newTreeNode(*pNum)
			currentTn.Right = newTn
			queue.add(newTn, currentTnd.depth+1)
		}
	}
	return root, nil
}

// getPNum 配列要素があれば値を返し、無ければ第二引数をfalseにして返却します。
func getPNum(pNums []*int, index int) (*int, bool) {
	if index > len(pNums)-1 {
		return nil, false
	}
	return pNums[index], true
}

package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/

func countComponents(n int, edges [][]int) int {
	searcher := NewGraphSearcher(n, edges)
	return searcher.SearchComponentCount()
}

type GraphSearcher struct {
	graph          map[int][]int       // key=node番号、value=keyのノードとつながっているnode番号
	visited        map[int]interface{} // 訪れたnode番号を保持
	componentCount int
}

func NewGraphSearcher(nodeNum int, edges [][]int) *GraphSearcher {
	graph := convertGraph(nodeNum, edges)
	return &GraphSearcher{
		graph:   graph,
		visited: make(map[int]interface{}),
	}
}

// convertGraph ノード数とエッジからgraphを作成します。
func convertGraph(nodeNum int, edges [][]int) map[int][]int {
	// nodeNum分のノードを作成
	graph := make(map[int][]int)
	for i := 0; i < nodeNum; i++ {
		graph[i] = []int{}
	}

	// ノードにつながっているノードを接続
	for _, edge := range edges {
		sNode := edge[0]
		dNode := edge[1]
		graph[sNode] = append(graph[sNode], dNode)
		graph[dNode] = append(graph[dNode], sNode)
	}
	return graph
}

// SearchComponentCount グラフを探索して繋がりのあるコンポーネントの個数を返却します。
func (g *GraphSearcher) SearchComponentCount() int {
	for keyNode := range g.graph {
		if g.isVisited(keyNode) {
			continue
		}
		g.componentCount++
		g.visit(keyNode)

		// 繋がっているノードに再帰的に訪れる
		g.searchNode(keyNode, 1)
	}
	return g.componentCount
}

func (g *GraphSearcher) visit(node int) {
	g.visited[node] = new(interface{})
}

func (g *GraphSearcher) isVisited(node int) bool {
	_, ok := g.visited[node]
	return ok
}

// search 繋がっているノードに再帰的に訪れる
func (g *GraphSearcher) searchNode(node int, depth int) {
	// fmt.Printf("node:%d, depth:%d\n", node, depth)
	g.visit(node)
	nodes, ok := g.graph[node]
	if !ok {
		return
	}

	for _, node := range nodes {
		if g.isVisited(node) {
			continue
		}
		g.searchNode(node, depth+1)
	}
}

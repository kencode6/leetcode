package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	searcher := NewPathSearcher(obstacleGrid)
	return searcher.Search()
}

type PathSearcher struct {
	obstacleGrid [][]int
	pathNumGrid  [][]int
	maxPoint     *Point
}

func NewPathSearcher(obstacleGrid [][]int) *PathSearcher {
	maxRow := len(obstacleGrid) - 1
	maxCol := len(obstacleGrid[0]) - 1

	maxPoint := NewPoint(maxRow, maxCol)

	pathNumGrid := createPathNumGrid(obstacleGrid)

	return &PathSearcher{
		obstacleGrid: obstacleGrid,
		pathNumGrid:  pathNumGrid,
		maxPoint:     maxPoint,
	}
}

func createPathNumGrid(obstacleGrid [][]int) [][]int {
	// パス数格納配列を複製する。要素は-1で初期化する。
	pathNumGrid := [][]int{}
	for i := 0; i < len(obstacleGrid); i++ {
		row := []int{}
		for j := 0; j < len(obstacleGrid[i]); j++ {
			row = append(row, -1)
		}
		pathNumGrid = append(pathNumGrid, row)
	}
	return pathNumGrid
}

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func p(x int, y int) *Point {
	return NewPoint(x, y)
}

func (s *PathSearcher) isBlock(p *Point) bool {
	return s.obstacleGrid[p.x][p.y] == 1
}

func (s *PathSearcher) pathNum(p *Point) int {
	return s.pathNumGrid[p.x][p.y]
}

func (s *PathSearcher) setPathNum(p *Point, pathNum int) {
	s.pathNumGrid[p.x][p.y] = pathNum
}

func (s *PathSearcher) maxPathNum() int {
	return s.pathNum(s.maxPoint)
}

func (s *PathSearcher) Search() int {
	// 初回がブロックの場合、終了
	if s.isBlock(p(0, 0)) {
		return 0
	}

	// 左と上の数値を加算するというロジックにする
	s.setPathNum(p(0, 0), 1)

	if s.maxPoint.x > 0 {
		// 初回の行
		for i := 1; i <= s.maxPoint.x; i++ {
			if s.isBlock(p(i, 0)) {
				// 岩を見つけたら対象ブロックを0に
				s.setPathNum(p(i, 0), 0)
			} else {
				// 岩でない場合は左と同じ数値
				prevPathNum := s.pathNum(p(i-1, 0))
				s.setPathNum(p(i, 0), prevPathNum)
			}
		}
	}

	if s.maxPoint.y > 0 {
		// 初回の列
		for i := 1; i <= s.maxPoint.y; i++ {
			if s.isBlock(p(0, i)) {
				// 岩を見つけたら対象ブロックを0に
				s.setPathNum(p(0, i), 0)
			} else {
				// 岩でない場合は上と同じ数値
				prevPathNum := s.pathNum(p(0, i-1))
				s.setPathNum(p(0, i), prevPathNum)
			}
		}
	}

	// gridが1行または1列の場合
	if s.maxPoint.x == 0 || s.maxPoint.y == 0 {
		return s.maxPathNum()
	}

	// 2回目以降の行
	for i := 1; i <= s.maxPoint.x; i++ {
		for j := 1; j <= s.maxPoint.y; j++ {
			if s.isBlock(p(i, j)) {
				// 岩を見つけたら対象ブロックを0に
				s.setPathNum(p(i, j), 0)
			} else {
				// 岩以外は左と上の数値を加算
				leftPathNum := s.pathNum(p(i-1, j))
				topPathNum := s.pathNum(p(i, j-1))
				pathNum := leftPathNum + topPathNum
				s.setPathNum(p(i, j), pathNum)
			}
		}
	}
	return s.maxPathNum()
}

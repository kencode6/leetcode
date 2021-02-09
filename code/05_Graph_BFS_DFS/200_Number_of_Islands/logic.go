package logic

import "fmt"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	searcher := NewIslandSeacher(grid)
	for {
		// 拠点を探索、拠点が見つからなかったら終了
		isFind := searcher.findPoint()
		if !isFind {
			break
		}
		// 島を探索する
		searcher.walkIsland()
	}
	return searcher.IslandNum()
}

type Point struct {
	row int
	col int
}

func NewPoint(row int, col int) *Point {
	return &Point{row: row, col: col}
}

type IslandSeacher struct {
	grid       [][]byte
	islandNum  int
	visitedMap map[string]bool // key=row-col, val=訪れたフラグ
	nowPoint   *Point

	rowSize int
	colSize int
}

func NewIslandSeacher(grid [][]byte) *IslandSeacher {

	rowSize := len(grid)
	colSize := len(grid[0])

	return &IslandSeacher{
		grid:       grid,
		visitedMap: make(map[string]bool),
		rowSize:    rowSize,
		colSize:    colSize,
	}
}

func (s *IslandSeacher) IslandNum() int {
	return s.islandNum
}

func (s *IslandSeacher) findPoint() bool {
	for row, vals := range s.grid {
		for col := range vals {
			// 訪れていないかつ島である場合、拠点にする。
			point := NewPoint(row, col)
			if !s.isVisited(point) && s.isIsland(point) {
				fmt.Printf("拠点:%d,%d islandNum:%d \n", row, col, s.islandNum)
				s.nowPoint = point
				s.islandNum++
				return true
			}

		}
	}
	return false
}

func (s *IslandSeacher) isVisited(p *Point) bool {
	visitedKey := visitedKey(p.row, p.col)
	return s.visitedMap[visitedKey]
}

func (s *IslandSeacher) isIsland(p *Point) bool {
	return string(s.grid[p.row][p.col]) == "1"
}

func (s *IslandSeacher) setVisited(p *Point) {
	fmt.Printf("訪れた:%d,%d islandNum:%d \n", p.row, p.col, s.islandNum)
	visitedKey := visitedKey(p.row, p.col)
	s.visitedMap[visitedKey] = true
}

func visitedKey(row, col int) string {
	return fmt.Sprintf("%d-%d", row, col)
}

func (s *IslandSeacher) walkIsland() {
	s.walk(s.nowPoint)
}

func (s *IslandSeacher) walk(point *Point) {

	// 訪れる
	s.setVisited(point)

	// 上下左右の探索ポイント作成
	points := []*Point{}

	// 上
	if point.row-1 >= 0 {
		points = append(points, NewPoint(point.row-1, point.col))
	}

	// 左
	if point.col-1 >= 0 {
		points = append(points, NewPoint(point.row, point.col-1))
	}

	// 下
	if point.row+1 < s.rowSize {
		points = append(points, NewPoint(point.row+1, point.col))
	}

	// 右
	if point.col+1 < s.colSize {
		points = append(points, NewPoint(point.row, point.col+1))
	}

	// 上下左右を探索
	for _, p := range points {
		//　島で無い、または訪れている場合スキップ
		if !s.isIsland(p) || s.isVisited(p) {
			continue
		}
		// 探索
		s.walk(p)
	}
}

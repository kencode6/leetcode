package logic

import "fmt"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 初回がブロックの場合、終了
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// 左と上の数値を再帰的に取得し求める方法
	// S(0,0) = 1
	// S(x,y) = 0 (x < 0, y < 0)
	// S(x,y) = S(x-1,y) + S(x,y-1)
	counter := NewPathCounter(obstacleGrid)
	count := counter.PathCount()
	fmt.Printf("count:%d", count)
	return count
}

type PathCounter struct {
	obstacleGrid [][]int
	cache        map[string]int
	maxX         int
	maxY         int
	calcCount    int
}

func NewPathCounter(obstacleGrid [][]int) *PathCounter {
	maxX := len(obstacleGrid) - 1
	maxY := len(obstacleGrid[0]) - 1

	return &PathCounter{
		obstacleGrid: obstacleGrid,
		cache:        make(map[string]int),
		maxX:         maxX,
		maxY:         maxY,
		calcCount:    0,
	}
}

func (p *PathCounter) pathCount(x int, y int, depth int) int {
	key := fmt.Sprintf("%d-%d", x, y)
	if count, ok := p.cache[key]; ok {
		return count
	}

	p.calcCount++

	// 開始点
	if x == 0 && y == 0 {
		return 1
	}

	// 範囲外の場合
	if x < 0 || y < 0 || x > p.maxX || y > p.maxY {
		return 0
	}

	// 岩の場合
	if p.obstacleGrid[x][y] == 1 {
		return 0
	}

	depth++

	// 左と上の数値を再帰的に取得
	leftCount := p.pathCount(x-1, y, depth)
	topCount := p.pathCount(x, y-1, depth)
	count := leftCount + topCount

	p.cache[key] = count
	return count
}

func (p *PathCounter) PathCount() int {
	return p.pathCount(p.maxX, p.maxY, 1)
}

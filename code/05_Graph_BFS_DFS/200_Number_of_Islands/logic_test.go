package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestNumIslands(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		InputConverter(
			func(index int, input interface{}) interface{} {
				return toByteSlices(input.([][]string))
			},
		).
		Executor(
			func(inputs []interface{}) interface{} {
				return numIslands(inputs[0].([][]byte))
			},
		).Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

// ToByteSlices [][]string を[][]byte に変換します。(05_Graph_BFS_DFSの島を作るための関数)
func toByteSlices(strSlices [][]string) [][]byte {
	byteSlices := [][]byte{}
	for _, strs := range strSlices {
		bytes := []byte{}
		for _, str := range strs {
			b := []byte(str)
			bytes = append(bytes, b[0])
		}
		byteSlices = append(byteSlices, bytes)
	}
	return byteSlices
}

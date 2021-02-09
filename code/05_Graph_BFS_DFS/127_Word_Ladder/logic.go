package logic

import "fmt"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	searcher := NewGraphSearcher(wordList)
	return searcher.SearchMinDistance(beginWord, endWord)
}

type GraphSearcher struct {
	graph   map[string][]string    // key=一文字を*にしたwildcardの単語、value=繋がりのある単語
	visited map[string]interface{} // 訪れた単語を保持
}

func NewGraphSearcher(wordList []string) *GraphSearcher {
	graph := convertGraph(wordList)
	return &GraphSearcher{
		graph:   graph,
		visited: make(map[string]interface{}),
	}
}

// WordData 単語と距離を保持
type WordData struct {
	word     string
	distance int
}

func NewWordData(word string, distance int) *WordData {
	return &WordData{
		word:     word,
		distance: distance,
	}
}

// WordDataQueue WordDataを格納するキュー
type WordDataQueue struct {
	wordDatas []*WordData
}

func NewWordDataQueue() *WordDataQueue {
	return &WordDataQueue{}
}

// Add 単語と距離をWordDataとして登録する
func (w *WordDataQueue) Add(word string, distance int) {
	data := NewWordData(word, distance)
	w.wordDatas = append(w.wordDatas, data)
}

// Poll 先頭のデータを取り出して破棄する。データがない場合なnilを返却する
func (w *WordDataQueue) Poll() *WordData {
	if len(w.wordDatas) == 0 {
		return nil
	}
	data := w.wordDatas[0]
	w.wordDatas = w.wordDatas[1:]
	return data
}

// convertGraph 単語リストからgraphを作成します。
func convertGraph(wordList []string) map[string][]string {
	graph := make(map[string][]string)
	for _, word := range wordList {

		// "hot" → ["*ot", "h*t", "*ot"]のワイルドカードを作成
		wcWords := convertWirdCardWords(word)

		// ワイルドカードをキーに単語を登録
		// 例えば"*ot"のようなワイルドカードのslice値は["hot", "dot", "lot"]になる。
		for _, wcWord := range wcWords {
			words, ok := graph[wcWord]
			if !ok {
				words = []string{}
			}
			words = append(words, word)
			graph[wcWord] = words
		}
	}
	return graph
}

// convertWirdCardWords wordを一文字ワイルドカードにした文字列のsliceを返却します。
// 例)
// "hot" → ["*ot", "h*t", "*ot"]
func convertWirdCardWords(word string) []string {
	wcWords := []string{}
	for i := 0; i < len(word); i++ {
		wcWord := word[:i] + "*" + word[i+1:]
		wcWords = append(wcWords, wcWord)
	}
	return wcWords
}

// SearchMinDistance 単語グラフを探索してbeginWordからendWordまでの最短距離を返却します。
func (g *GraphSearcher) SearchMinDistance(beginWord string, endWord string) int {
	// queueを生成し、初回の単語を登録
	queue := NewWordDataQueue()
	queue.Add(beginWord, 1)

	for {
		wordData := queue.Poll()
		if wordData == nil {
			// endwordが見つからない
			break
		}

		fmt.Printf("word:%s depth:%d\n", wordData.word, wordData.distance)

		// ワイルドカードを取得
		wcWords := convertWirdCardWords(wordData.word)
		for _, wcWord := range wcWords {
			//ワイルドカードと繋がりのある単語を取得
			words, ok := g.graph[wcWord]
			if !ok {
				continue
			}

			fmt.Printf("wcWord:%s, words:%v, depth:%d\n", wcWord, words, wordData.distance)

			for _, word := range words {
				// すでに探索済みの場合はスキップ
				if g.isVisited(word) {
					continue
				}

				if word == endWord {
					// endWordを発見
					return wordData.distance + 1
				}
				// 探索済みにしてキューに登録
				g.visit(word)
				queue.Add(word, wordData.distance+1)
			}
		}
	}
	return 0
}

func (g *GraphSearcher) visit(word string) {
	g.visited[word] = new(interface{})
}

func (g *GraphSearcher) isVisited(word string) bool {
	_, ok := g.visited[word]
	return ok
}

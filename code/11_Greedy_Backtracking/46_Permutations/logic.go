package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	builder := NewPermuteBuilder(nums)
	builder.Build()
	return builder.permuteNumsSlice
}

type PermuteBuilder struct {
	nums             []int
	permuteNumsSlice [][]int
}

func NewPermuteBuilder(nums []int) *PermuteBuilder {
	return &PermuteBuilder{
		nums:             nums,
		permuteNumsSlice: [][]int{},
	}
}

func (b *PermuteBuilder) Build() {
	b.build([]int{}, b.nums)
}

// build 選択数値:selectedNumsと選択対象数値:remainingNumsを元に再帰的に組合せを登録します
func (b *PermuteBuilder) build(selectedNums []int, remainingNums []int) {
	if len(selectedNums) == len(b.nums) {
		// 終端
		b.permuteNumsSlice = append(b.permuteNumsSlice, selectedNums)
		return
	}

	for i := 0; i < len(remainingNums); i++ {
		num := remainingNums[i]
		// 選択数値にnumを追加
		newSelectedNums := newAppendNums(selectedNums, num)

		// 選択対象数値から指定indexを除外
		newRemainingNums := newRemoveNums(remainingNums, i)

		// 再帰呼び出し
		b.build(newSelectedNums, newRemainingNums)
	}
}

// newAppendNums numsにappendNumを追加したsliceを生成します
func newAppendNums(nums []int, appendNum int) []int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	newNums = append(newNums, appendNum)
	return newNums
}

// newRemoveNums numsからremoveIndexを取り除いたsliceを生成します
func newRemoveNums(nums []int, removeIndex int) []int {
	if len(nums) == 1 {
		return []int{}
	}
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	newNums = append(newNums[:removeIndex], newNums[removeIndex+1:]...)
	return newNums
}

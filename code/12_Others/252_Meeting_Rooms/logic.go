package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/meeting-rooms/

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) <= 1 {
		return true
	}

	for i := 1; i < len(intervals); i++ {
		// 新規meetingチェック対象の時間
		checkInterval := intervals[i]

		for j := 0; j < i; j++ {
			// 既存meeting確認時間
			interval := intervals[j]
			if isDuplicate(checkInterval, interval) {
				return false
			}
		}
	}
	return true
}

// intervalが重なっているかを確認
func isDuplicate(interval []int, newInterval []int) bool {
	start := interval[0]
	end := interval[1]
	newStart := newInterval[0]
	newEnd := newInterval[1]

	if start <= newStart && newStart < end {
		// 新規meetingの開始時間が既存meetingの終了時間より前
		return true
	}

	if newEnd <= end && start < newEnd {
		// 新規meetingの終了時間が既存meetingの開始時間より前
		return true
	}

	if newStart <= start && end <= newEnd {
		// 新規meetingの開始時間、終了時間が既存meetingの時間に含まれている
		return true
	}
	return false
}

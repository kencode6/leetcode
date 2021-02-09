package logic

import "sort"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/meeting-rooms-ii/

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	if len(intervals) == 1 {
		return 1
	}

	// 開始時間でソート
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 各部屋ごとにintervalsを格納できる複数のmeetingroomを用意
	meetingRooms := [][][]int{}
	for _, newInterval := range intervals {
		if len(meetingRooms) == 0 {
			// 初回
			meetingRooms = append(meetingRooms, [][]int{newInterval})
			continue
		}

		// 既存のmeetingRoomsにintervalsを入れられるかを確認
		insertRoomIndex, isInsert := getInsertRoomIndex(meetingRooms, newInterval)

		if isInsert {
			// 既存のmeetingroomにintervalsを追加
			meetingRooms[insertRoomIndex] = append(meetingRooms[insertRoomIndex], newInterval)
		} else {
			// 新しいmeeting roomを作成しintervalsを追加
			meetingRooms = append(meetingRooms, [][]int{newInterval})
		}
	}
	return len(meetingRooms)
}

// getInsertRoomIndex 既存のmeetingroomに予定が入れられるかを確認する。
// 入れられればmeetingRoomのindex番号を返す。
// return (meetingRoomのindex番号, 登録可能であるか)
func getInsertRoomIndex(meetingRooms [][][]int, newInterval []int) (insertRoomIndex int, isEnableInsert bool) {
	for i := 0; i < len(meetingRooms); i++ {
		meetingRoom := meetingRooms[i]

		isInsert := true
		for _, interval := range meetingRoom {
			if isDuplicate(interval, newInterval) {
				isInsert = false
				break
			}
		}
		if isInsert {
			insertRoomIndex = i
			isEnableInsert = true
			return
		}
	}
	isEnableInsert = false
	return
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

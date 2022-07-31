# golang_meeting_rooms_v2

Given an array of meeting time intervals consisting of start and end times `[[s1,e1],[s2,e2],...] (si < ei)`
, find the minimum number of conference rooms required.)

## Examples

**Example1**

```
Input: intervals = [(0,30),(5,10),(15,20)]
Output: 2
Explanation:
We need two meeting rooms
room1: (0,30)
room2: (5,10),(15,20)

```

**Example2**

```
Input: intervals = [(2,7)]
Output: 1
Explanation:
Only need one meeting room

```

## 解析

給定一個 2D 陣列 intervals 

每個 intervals[i] = [ $start_i$, $end_i$] 代表一個時間區間 $start_i$ ≤ time < $end_i$fu

假設遇到兩個時間區間重疊了，就必須要多開一間房間

要求寫一個演算法計算最少需要開幾間房間才能順利舉行所有會議

當會議開始則需要開啟一間房間

因為要開最少房間，所以會議結束會關閉一個房間

具體作法如下圖：

![](https://i.imgur.com/8JjDEPb.png)

一開始我們需要把所有 intervals 的開始與結束時間放在兩個陣列 start, end

然後把 start, end 由小至大做排序 

然後我們可以使用兩個指標 pStart： 代表 start 目前所在位置，pEnd: 代表 start 目前所在位置

初始化 pStart, pEnd := 0, 0, result:=0 , count := 0 

從上圖可以看出來 當 pStart 走到 n-1 就可以拿到最大值 因為代表不會再有會議需要開始

當 pStart < n , 檢查 start[pStart] < end[pEnd] 

如果是,代表在目前開始會議之前沒有會議要結束

所以需要把 count += 1, pStart += 1

如果否, 代表在目前開始會議之前有會議要結束

所以需要把 count -= 1, pEnd += 1

更新 result = max(result, count)

當 pStart 走到最後 回傳 result

因為 除了 sort 要花 O(nlogn) 以上演算法只要花 O(n) 時間複雜度

所以整個時間複雜度是 O(nlogn)

需要額外紀錄兩個陣列  start, end 來做排序

所以空間複雜度是 O(n)

## 程式碼
```go
package sol

import "sort"

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */
/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */
func MinMeetingRooms(intervals []*Interval) int {
	count, result := 0, 0
	n := len(intervals)
	start, end := make([]int, n), make([]int, n)
	for pos := 0; pos < n; pos++ {
		start[pos] = intervals[pos].Start
		end[pos] = intervals[pos].End
	}
	sort.Ints(start)
	sort.Ints(end)
	pStart, pEnd := 0, 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for pStart < n {
		if start[pStart] < end[pEnd] {
			pStart++
			count++
		} else {
			pEnd++
			count--
		}
		result = max(result, count)
	}
	return result
}

```
## 困難點

1. 需要想出如何從開始與結束時間與開啟房間與關閉房間的關係
2. 理解 two pointer 來解決這個問題

## Solve Point

- [x]  建立兩個陣列 start, end來 儲存所有 intervals 內的開始與結束時間
- [x]  對 start, end 由小到大排序
- [x]  初始化 pStart, pEnd := 0, 0, result:=0 , count := 0
- [x]  當 pStart < n 檢查 start[pStart] < end[pEnd]
- [x]  如果是,代表在目前開始會議之前沒有會議要結束, 更新 count += 1, pStart += 1
- [x]  如果否, 代表在目前開始會議之前有會議要結束, 更新 count -= 1, pEnd += 1
- [x]  更新 result = max(result, count)
- [x]  當 pStart 走到最後 回傳 result
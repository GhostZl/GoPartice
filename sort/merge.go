package sort

func MergeSort(arr []int) {
	split(arr, 0, len(arr)-1)
}

func split(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1
	split(arr, start, mid)
	split(arr, mid+1, end)
	merge(arr, start, end, mid)
}

func merge(arr []int, start int, end int, mid int) {
	temp := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		temp[i-start] = arr[i]
	}
	startL := 0
	startR := mid + 1 - start
	tempMid := mid - start
	tempEnd := len(temp) - 1
	for i := 0; i < len(temp); i++ {
		if startL > tempMid {
			arr[start+i] = temp[startR]
			startR++
			continue
		}
		if startR > tempEnd {
			arr[start+i] = temp[startL]
			startL++
			continue
		}
		if temp[startL] <= temp[startR] {
			arr[start+i] = temp[startL]
			startL++
			continue
		}
		arr[start+i] = temp[startR]
		startR++
	}
}

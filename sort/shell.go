package sort

//插入排序的优化
func ShellSort(arr []int) {
	gap := len(arr) >> 1
	for {
		for start := 0; start <= gap; start++ {
			addCount := gap + 1
			for i := start; i < len(arr)-addCount; i += addCount {
				j := i + addCount
				for j-addCount >= 0 && arr[j] < arr[j-addCount] {
					arr[j], arr[j-addCount] = arr[j-addCount], arr[j]
					j -= addCount
				}
			}
		}
		if gap == 0 {
			break
		}
		gap = gap >> 1
	}
}

package sort

func InserSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		j := i + 1
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}

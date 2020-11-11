package sort

import "fmt"

func BubbleSort(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(count)
	return arr
}

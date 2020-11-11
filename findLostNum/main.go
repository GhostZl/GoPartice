package main

import "fmt"

func findLostNum(nums []int) int {
	var maxNum int = 0
	var minNum int = 0
	var expected int = 0
	var real int = 0
	for i, num := range nums {
		if i==0 {
			minNum = num
			maxNum = num
		}
		if maxNum < num {
			maxNum = num
		}
		if minNum > num {
			minNum = num
		}
		real += num
	}
	length:=maxNum-minNum+1
	expected = minNum*length + (length-1)*length/2
	return expected-real
}
func main() {
	fmt.Println(findLostNum([]int{2,4,5,3,1,7}))
}

package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := [][]int{
		[]int{2, 31, 4, 1, 5, 3},
	}
	for _, ca := range cases {
		BubbleSort(ca)
		fmt.Println(ca)
	}
}

func TestMergeSort(t *testing.T) {
	cases := [][]int{
		[]int{2, 31, 4, 4, 1, 5, 3},
		[]int{2, 1},
	}
	for _, ca := range cases {
		MergeSort(ca)
		fmt.Println(ca)
	}
}

func TestInserSort(t *testing.T) {
	cases := [][]int{
		[]int{2, 31, 4, 4, 1, 5, 3},
	}
	for _, ca := range cases {
		InserSort(ca)
		fmt.Println(ca)
	}
}

func TestShellSort(t *testing.T) {
	cases := [][]int{
		[]int{2, 31, 4, 4, 1, 5, 3},
		[]int{},
	}
	for _, ca := range cases {
		ShellSort(ca)
		fmt.Println(ca)
	}
}

func TestSelectSort(t *testing.T) {
	cases := [][]int{
		[]int{2, 31, 4, 4, 1, 5, 3},
		[]int{},
	}
	for _, ca := range cases {
		SelectSort(ca)
		fmt.Println(ca)
	}
}

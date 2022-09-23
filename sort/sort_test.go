package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{10, 3, 1, 3, 5, 8}
	fmt.Println(SelectionSort(nums))
}

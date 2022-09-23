package search

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	if target != BinarySearch(nums, target) {
		t.Failed()
	}
}

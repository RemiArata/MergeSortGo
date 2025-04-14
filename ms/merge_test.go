package ms

import (
	"MergeSortGo/helpers"
	"testing"
)

func TestSortedLenTwo(t *testing.T) {
	data := []int{1, 2}
	s := MergeSort(data)

	if s[0] > s[1] {
		t.Error("List was not sorted correctly")
	}
}

func TestUnsortedLenTwo(t *testing.T) {
	data := []int{2, 0}
	s := MergeSort(data)

	if s[0] > s[1] {
		t.Error("List was not sorted correctly")
	}
}

func TestSortedLenFive(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := MergeSort(data)

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Error("List has elements unordered", s[i-1], "isn't less than", s[i])
		}
	}
}

func TestUnsortedLenFive(t *testing.T) {
	data := []int{5, 1, 3, 2, 4}
	s := MergeSort(data)

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Error("List has elements unordered", s[i-1], "isn't less than", s[i])
		}
	}
}

func TestUnsortedRandomOne(t *testing.T) {
	data := helpers.GenerateData(false, 10)
	s := MergeSort(data)

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Error("List has elements unordered", s[i-1], "isn't less than", s[i])
		}
	}
}

func TestUnsortedRandomTwo(t *testing.T) {
	data := helpers.GenerateData(false, 50)
	s := MergeSort(data)

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Error("List has elements unordered", s[i-1], "isn't less than", s[i])
		}
	}
}

func TestUnsortedRandomThree(t *testing.T) {
	data := helpers.GenerateData(false, 100)
	s := MergeSort(data)

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Error("List has elements unordered", s[i-1], "isn't less than", s[i])
		}
	}
}

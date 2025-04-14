package ms

func MergeSort(data []int) []int {
	if len(data) == 1 {
		return data
	} else if len(data) <= 2 {
		if data[0] > data[1] {
			return []int{data[1], data[0]}
		} else {
			return []int{data[0], data[1]}
		}
	}
	mid := len(data) / 2
	s1 := MergeSort(data[:mid])
	s2 := MergeSort(data[mid:])

	i, j := 0, 0
	newSlice := []int{}

	for i+j < len(s1)+len(s2) {
		if i < len(s1) && j < len(s2) {
			if s1[i] < s2[j] {
				newSlice = append(newSlice, s1[i])
				i = i + 1
			} else {
				newSlice = append(newSlice, s2[j])
				j = j + 1
			}
		} else if i < len(s1) {
			newSlice = append(newSlice, s1[i])
			i = i + 1
		} else if j < len(s2) {
			newSlice = append(newSlice, s2[j])
			j = j + 1
		}
	}
	return newSlice
}

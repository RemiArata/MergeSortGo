package helpers

import "math/rand"

func GenerateData(useSeed bool, length int) []int {
	if useSeed {
		rand.Seed(5)
	}
	var d []int
	for i := 0; i < length; i++ {
		d = append(d, rand.Intn(100))
	}
	return d
}

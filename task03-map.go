package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	return sort.Ints(keys)
}

package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, len(input))
	output := make([]string, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		//fmt.Println(k)
		i++
	}
	sort.Ints(keys)
	for k := range keys {
		output[k] = input[keys[k]]
	}

	return output
}


package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for i := 0; i < len(input)/2; i++ {
		j := len(input) - i - 1
		input[i], input[j] = input[j], input[i]
	}
	return input

}

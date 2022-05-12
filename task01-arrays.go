package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32 = 0.0
	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}
	var length float32 = float32(len(input))
	var ans float32 = sum / length
	return ans
}

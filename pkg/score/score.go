package score

func Score(numbers []int) int {
	left := 10*numbers[0] + numbers[1]
	right := 10*numbers[2] + numbers[3]
	return left + right
}

func MaxScore(nns [][]int) int {
	max := 0
	for _, ns := range nns {
		score := Score(ns)
		if score > max {
			max = score
		}
	}
	return max
}

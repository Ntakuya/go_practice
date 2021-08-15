package countpositivessumnegatives

func CountPositivesSumNegatives(numbers []int)[]int {
	res := []int{0, 0}

	for _, num := range numbers {
		if (num < 0) {
			res[1] += num
		} else {
			res[0]++
		}
	}

	return res
}

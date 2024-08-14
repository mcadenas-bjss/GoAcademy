package arrayandslices

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		t := 0
		if len(numbers) > 0 {
			t = Sum(numbers[1:])
		}
		sums = append(sums, t)
	}

	return sums
}

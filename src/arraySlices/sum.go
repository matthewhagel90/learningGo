package arraySlices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	var sum []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
			continue
		}
		tail := numbers[1:]
		sum = append(sum, Sum(tail))
	}
	return sum
}

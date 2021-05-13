package main

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return sums
}

func main() {
}

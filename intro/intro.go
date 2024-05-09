package main

func Hello(name string) string {
	return "Hello, " + name
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, slc := range numbers {
		sums = append(sums, Sum(slc))
	}
	return sums
}

package listStuff

import "sort"

func Array1To10() [10]int {
	return [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func SortIntArray(arr [10]int, direction int) [10]int {
	sort.SliceStable(arr[:], func(i, j int) bool {
		return (arr[i]-arr[j])*direction < 0
	})

	return arr
}

func SortEvenAscOddDesc(arr []int) ([]int, []int) {
	odd := make([]int, 0)
	even := make([]int, 0)

	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	sort.Ints(even)
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	return even, odd
}

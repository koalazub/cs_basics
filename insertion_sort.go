package main

// straight from the Introduction to Algorithms chapt. 2 section on Insertion sort
func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {

		k := list[i]
		j := i

		for j > 0 && list[j-1] > k {
			list[j+1] = list[j-1]
			j -= 1
			list[j] = k
		}
	}

	return list

}

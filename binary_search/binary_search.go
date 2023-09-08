package binarysearch

func BinarySearch(numbers []int, target int) int {
	L := 0
	R := len(numbers) - 1
	for L <= R {
		m := (L + R) / 2

		if numbers[m] == target {
			return numbers[m]
		}
		if numbers[m] < target {
			L = m + 1
		}

		if numbers[m] > target {
			R = m - 1
		}
	}
	return 0
}

package blind75

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, 0)
	for i := len(bucket) - 1; len(res) != k; i-- {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
		}
	}

	return res
}

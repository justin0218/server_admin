package tool

func IntInArray(item int, arr []int) (dx int) {
	for idx, v := range arr {
		if v == item {
			dx = idx
			return
		}
	}
	return -1
}

package tantan

func MergeArr(arrs ...[]int) []int {
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		return arrs[0]
	}
	res := arrs[0]
	for i := 1; i < len(arrs); i++ {
		indexA := 0
		indexB := 0
		temp := []int{}
		for indexA < len(arrs[i]) || indexB < len(res) {
			if indexA >= len(arrs[i]) && indexB < len(res) {
				temp = append(temp, res[indexB])
				indexB++
				continue
			}
			if indexB >= len(res) && indexA < len(arrs[i]) {
				temp = append(temp, arrs[i][indexA])
				indexA++
				continue
			}
			if arrs[i][indexA] < res[indexB] {
				temp = append(temp, arrs[i][indexA])
				indexA++
			} else {
				temp = append(temp, res[indexB])
				indexB++
			}
		}
		res = temp
	}
	return res
}

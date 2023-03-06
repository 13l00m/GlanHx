package utils

import "strconv"

func RemoveDuplicate(elements []int) []int {
	// 使用 map 去重
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			continue
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func RemoveDuplicate_String(elements []string) []string {
	result := make([]string, 0)

	for _, v := range elements {
		if !contains(result, v) {
			result = append(result, v)
		}
	}
	return result

}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
func StringArray2IntArray(s []string) []int {
	int_arr := make([]int, len(s))

	for i, p_str := range s {
		p_int, err := strconv.Atoi(p_str)
		if err != nil {
			return []int{80, 443}
		}

		int_arr[i] = p_int
	}
	return int_arr
}

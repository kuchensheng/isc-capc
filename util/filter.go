package util

//ListFilter 对列表进行过滤
func ListFilter[T any](list []T, filter func(item T) bool) []T {
	var result []T
	if len(list) == 0 {
		return result
	}
	for _, t := range list {
		if filter(t) {
			result = append(result, t)
		}
	}
	return result
}

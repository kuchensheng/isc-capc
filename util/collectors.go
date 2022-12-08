package util

import "fmt"

func Map[T, R any](slice []T, handler func(T) R) []R {
	var result []R
	for _, t := range slice {
		func() {
			defer func() {
				if x := recover(); x != nil {
					fmt.Println("转换失败", x)
				}
			}()
			result = append(result, handler(t))
		}()
	}
	return result
}

package util

import (
	"reflect"
)

func getAllFieldsAndValues(obj any) map[string]any {
	of := reflect.ValueOf(obj)
	typeOf := of.Type()
	var result = make(map[string]any)
	for i := 0; i < of.NumField(); i++ {
		f := of.Field(i)
		result[typeOf.Field(i).Name] = f.Interface()
	}
	return result
}

func GetAllFields(obj any) []string {
	var result []string
	values := getAllFieldsAndValues(obj)
	for s, _ := range values {
		result = append(result, s)
	}
	return result
}

func GetNotNullFields(obj any) []string {
	var result []string
	values := getAllFieldsAndValues(obj)
	for s, a := range values {
		if a != nil {
			result = append(result, s)
		}
	}
	return result
}

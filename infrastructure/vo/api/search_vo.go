package api

type SearchVO struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Method      []string `json:"method"`
	Code        string   `json:"code"`
	CategoryIds []int    `json:"categoryIds"`
}

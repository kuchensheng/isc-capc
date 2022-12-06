package api

type SearchVO struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Method      []string `json:"method"`
	CategoryIds []int    `json:"categoryIds"`
}

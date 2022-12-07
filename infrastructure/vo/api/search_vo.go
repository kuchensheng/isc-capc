package api

type SearchVO struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Method      []string `json:"method"`
	Code        string   `json:"code"`
	CategoryIds []int    `json:"categoryIds"`
}

type DetailVO struct {
	Code    string `json:"code"`
	ID      int    `json:"id"`
	Path    string `json:"path"`
	Method  string `json:"method"`
	Version string `json:"version"`
}

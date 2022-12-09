package api

type SearchVO struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Methods     []string `json:"methods"`
	Codes       []string `json:"codes"`
	CategoryIds []int    `json:"categoryIds"`
	Ids         []int    `json:"ids"`
	Types       []int    `json:"types"`
}

type DetailVO struct {
	Code    string `json:"code"`
	ID      int    `json:"id"`
	Path    string `json:"path"`
	Method  string `json:"method"`
	Version string `json:"version"`
}

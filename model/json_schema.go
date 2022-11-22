package model

type Schema struct {
	Objects []struct {
		Name   string `json:"name"`
		Fields []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"fields"`
	} `json:"objects"`
	Query []struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Object    string `json:"object"`
		Arguments []struct {
		} `json:"arguments"`
	} `json:"query"`
	Mutations []struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Object    string `json:"object"`
		Arguments []struct {
		} `json:"arguments"`
	} `json:"mutations"`
}
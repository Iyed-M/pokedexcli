package api

type locationAreas struct {
	Count    float64 `json:"count,omitempty"`
	Next     string  `json:"next,omitempty"`
	Previous string  `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

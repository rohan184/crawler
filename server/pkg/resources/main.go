package resources

type RequestBody struct {
	URL string `json:"url"`
}

type Insight struct {
	URL        string   `json:"url"`
	WordCount  int      `json:"wordCount"`
	WebLinks   []string `json:"webLinks"`
	MediaLinks []string `json:"images"`
}

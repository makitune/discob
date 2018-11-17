package search

// made by https://mholt.github.io/json-to-go/

// Custom Search Engine JSON API Response
type Response struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind        string `json:"kind"`
		Title       string `json:"title"`
		HTMLTitle   string `json:"htmlTitle"`
		Link        string `json:"link"`
		DisplayLink string `json:"displayLink"`
		Snippet     string `json:"snippet"`
		HTMLSnippet string `json:"htmlSnippet"`
		Mime        string `json:"mime"`
		Image       struct {
			ContextLink     string `json:"contextLink"`
			Height          int    `json:"height"`
			Width           int    `json:"width"`
			ByteSize        int    `json:"byteSize"`
			ThumbnailLink   string `json:"thumbnailLink"`
			ThumbnailHeight int    `json:"thumbnailHeight"`
			ThumbnailWidth  int    `json:"thumbnailWidth"`
		} `json:"image"`
	} `json:"items"`
}

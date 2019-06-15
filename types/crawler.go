package types

// Crawler gets item links for some query

type Query struct {
	Query string
	City  string
}

type Crawler interface {
	// GetResults - make a request for a given query and return
	// url of pages that matching the request
	GetResults(q Query) ([]string, error)
}

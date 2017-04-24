package urlparser

import (
	"errors"
	"net/url"
)

type QueryURL struct {
	originalURL string
	urls        []string
}

//New creates a new instance of the QueryURL filtering out the invalid URLs
func New(rawURL, queryParam string) (*QueryURL, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, err
	}

	if queryParam == "" {
		return nil, errors.New("query parameter is empty")
	}

	queryURLS := parsedURL.Query()[queryParam]
	validURLs := []string{}
	for _, u := range queryURLS {
		if err := valid(u); err == nil {
			validURLs = append(validURLs, u)
		}
	}

	// fmt.Printf("queryparam: %v: %v\n", validURLs, parsedURL.Query()["u"])
	return &QueryURL{
		originalURL: rawURL,
		urls:        validURLs,
	}, nil
}

//AllURL return all the valid urls that have been stored
func (qu *QueryURL) AllURL() []string {
	return qu.urls
}

// Length gives the number of valid URLs
func (qu *QueryURL) Length() int {
	return len(qu.urls)
}

//extracts the 'u' query parameters ignores others
func valid(rawurl string) error {
	_, err := url.ParseRequestURI(rawurl)
	return err
}

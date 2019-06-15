package avito

import (
	"findthing/types"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
)

// Avito.ru crawler

const (
	ENDPOINT = "https://www.avito.ru"
)

type Crawler struct {

}

func makeURL(q types.Query) string {
	params := url.Values{}

	if len(q.Query) > 0 {
		params.Add("q", q.Query)
	}


	return fmt.Sprintf("%v/%v?%v", ENDPOINT, q.City, params.Encode())
}

func (c *Crawler) GetResults(q types.Query) ([]string, error) {
	urlString := makeURL(q)

	res, err := http.Get(urlString)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer res.Body.Close()

	page, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil
	}

	var results []string
	page.Find("a.item-description-title-link").Each(func(num int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			results = append(results, fmt.Sprintf("%v%v", ENDPOINT, href))
		}
	})

	return results, nil
}



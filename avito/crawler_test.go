package avito

import (
	"findthing/types"
	"log"
	"testing"
)

func TestCrawler_GetResults(t *testing.T) {
	c := Crawler{}

	q := types.Query{Query: "Таненбаум", City: "sankt-peterburg"}
	res, err := c.GetResults(q)
	if err != nil {
		t.Errorf("Failed to make request for '%v' in city '%v': %v", q.Query, q.City, err)
	}

	log.Println(res)
}

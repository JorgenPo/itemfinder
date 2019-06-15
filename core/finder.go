package core

import (
	"findthing/avito"
	"findthing/types"
	"log"
	time2 "time"
)

// Finder - finds items by query
type Finder struct {
	crawlers map[string]types.Crawler
	parsers map[string]types.Parser
}

func NewFinder() *Finder {
	crawlers := make(map[string]types.Crawler)
	parsers := make(map[string]types.Parser)

	crawlers["avito"] = &avito.Crawler{}
	parsers["avito"] = &avito.Parser{}

	return &Finder{crawlers, parsers}
}

func (f Finder) Find(q types.Query) (items []*types.Item) {
	log.Printf("Finding '%v'...", q)

	start := time2.Now()
	for name, crawler := range f.crawlers {
		log.Printf("Getting results in '%v'...", name)

		itemUrls, err := crawler.GetResults(q)
		if err != nil {
			log.Printf("Getting results failed: %v", err)
			continue
		}

		for number, url := range itemUrls {
			log.Printf("Parsing result #%v (%v)", number, url)

			item, err := f.parsers[name].Parse(url)
			if err != nil {
				log.Printf("Parsing results failed: %v", err)
				continue
			}

			items = append(items, item)
		}
	}

	log.Printf("Searching done. It takes '%v' seconds for '%v' results",
		time2.Now().Sub(start).Seconds(), len(items))

	return
}



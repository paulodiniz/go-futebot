package main

import (
	"github.com/gocolly/colly"
)

func FetchTable() map[string]string {
	var m map[string]string = make(map[string]string)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("www.gazetaesportiva.com"),
	)
	c.OnHTML(".table tr", func(e *colly.HTMLElement) {
		var team = e.ChildText("td.table__team")

		if len(team) > 0 {
			var points = e.ChildText("td.table__stats:nth-child(4)")
			m[team] = points
		}
	})

	c.Visit("https://www.gazetaesportiva.com/campeonatos/brasileiro-serie-a/")

	return m
}

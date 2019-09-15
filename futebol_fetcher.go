package main

import (
	"strconv"

	"github.com/gocolly/colly"
)

type Score struct {
	team   string
	points int
}

func FetchTable() []Score {
	var scores = []Score{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("www.gazetaesportiva.com"),
	)
	c.OnHTML(".table tr", func(e *colly.HTMLElement) {
		var team = e.ChildText("td.table__team")

		if len(team) > 0 {
			var pointsText = e.ChildText("td.table__stats:nth-child(4)")
			points, err := strconv.Atoi(pointsText)
			if err != nil {
				panic("Cannot convert")
			}

			scores = append(scores, Score{team: team, points: points})
		}
	})

	c.Visit("https://www.gazetaesportiva.com/campeonatos/brasileiro-serie-a/")

	return scores
}

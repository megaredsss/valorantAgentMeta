package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	var agentsSlice []string
	var mapsSlice []string
	var data map[string]map[string]int
	// var allMaps []map[string]map[string]int
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Connect:", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error in:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(".wf-card.mod-dark.mod-table.mod-scroll table.wf-table.mod-pr-global", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			// if i > 1 {
			// fmt.Println("0")
			// currentMap := make(map[string]map[string]int)

			tr.ForEach("th img", func(i int, th *colly.HTMLElement) {
				re := regexp.MustCompile("/img/vlr/game/agents/(.*).png")
				agentsSlice = append(agentsSlice, re.ReplaceAllString(th.Attr("src"), "${1}"))
				// fmt.Println(agentsSlice)
			})
			if i > 1 {
				tds := tr.DOM.Find("td")
				tds.Find("span").Remove()
				if tds.Length() > 0 {
					firstColumnText := tds.First().Text()
					firstColumnText = strings.TrimSpace(firstColumnText)
					mapsSlice = append(mapsSlice, firstColumnText)
				}
			}
		})
		// tr.ForEach("td", func(i int, td *colly.HTMLElement){
		// 				if i == 0{
		// 					currentMap[
		// 				}
		// }

		fmt.Println(mapsSlice)
		fmt.Println(agentsSlice)
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			tr.ForEach("td", func(j int, td *colly.HTMLElement) {
				// html, _ := td.DOM.Html()
				text := td.DOM.Find("tr.pr-global-row:nth-child(3) > td:nth-child(5) > div:nth-child(1)").Text()
				text = strings.TrimSpace(text)
				fmt.Println(text)
				// fmt.Printf("HTML ячейки %d в строке %d: %s\n", j, i, html)
				if i > 1 && j > 3 {
					data = make(map[string]map[string]int)
					data[mapsSlice[i-1]] = make(map[string]int)
					selection := td.DOM.Find("div.color-sq > span")
					fmt.Println(selection.Html())
					data[mapsSlice[i-2]][agentsSlice[j-4]], _ = strconv.Atoi(selection.Text())
					fmt.Println(data)
				}
			})
		})
		fmt.Println(data)
	})

	fmt.Println("Scrapping Completed")

	c.Visit("https://www.vlr.gg/event/agents/2097/valorant-champions-2024")
	// fmt.Println(data)
}

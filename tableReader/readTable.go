package tableReader

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

type Row map[string]string
type Table []Row

// ReadTable reading data from table on site
func ReadTable() Table {
	table := Table{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Connect:", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error in:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(".wf-card.mod-dark.mod-table.mod-scroll table.wf-table.mod-pr-global", func(e *colly.HTMLElement) {
		var headers []string
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			row := Row{}
			if i == 0 {
				tr.ForEach("th", func(i int, h *colly.HTMLElement) {
					if i == 0 {
						headers = append(headers, h.Text)
					}
				})
				tr.ForEach("th img", func(i int, th *colly.HTMLElement) {
					re := regexp.MustCompile("/img/vlr/game/agents/(.*).png")
					headers = append(headers, strings.TrimSpace(re.ReplaceAllString(th.Attr("src"), "${1}")))
				})
			} else if i > 1 {
				tr.ForEach("td", func(i int, td *colly.HTMLElement) {
					if td.Attr("style") == "white-space: nowrap; padding-top: 0; padding-bottom: 0;" {
						td.DOM.Find("span").Remove()
						if i < len(headers) {
							header := headers[i]
							row[header] = strings.TrimSpace(strings.TrimSpace(td.DOM.Text()))
						}
					}
					if td.Attr("class") == "mod-color-sq" {
						header := headers[i-3]
						row[header] = strings.TrimSpace(td.ChildText("span"))
					}
				})
				if len(row) > 0 {
					table = append(table, row)
				}
			}
		})
	})

	err := c.Visit("https://www.vlr.gg/event/agents/2097/valorant-champions-2024")
	if err != nil {
		fmt.Println("Error in Visit:", err)
	}
	return table
}

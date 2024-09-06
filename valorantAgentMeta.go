package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortMapByValues(m map[string]int) PairList {
	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))
	return p
}
func main() {

	type Row map[string]string
	// var allMaps []map[string]map[string]int
	type Table []Row
	// tables := []Table{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Connect:", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error in:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(".wf-card.mod-dark.mod-table.mod-scroll table.wf-table.mod-pr-global", func(e *colly.HTMLElement) {
		table := Table{}
		var headers []string
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			row := Row{}
			//fmt.Println(i)
			// if i > 1 {
			// fmt.Println("0")
			// currentMap := make(map[string]map[string]int)
			if i == 0 {
				tr.ForEach("th", func(i int, h *colly.HTMLElement) {
					if i == 0 {
						headers = append(headers, h.Text)
						//fmt.Println(h.Text)
					}
				})
				tr.ForEach("th img", func(i int, th *colly.HTMLElement) {
					re := regexp.MustCompile("/img/vlr/game/agents/(.*).png")
					headers = append(headers, strings.TrimSpace(re.ReplaceAllString(th.Attr("src"), "${1}")))
					//fmt.Println(headers)
				})
			} else if i > 1 {
				//fmt.Println(i)
				tr.ForEach("td", func(i int, td *colly.HTMLElement) {
					//fmt.Println("td = ", td)
					//fmt.Println("td.Class = ", td.Attr("class"))
					//fmt.Println("td.Attr = ", td.Attr("style"))
					if td.Attr("style") == "white-space: nowrap; padding-top: 0; padding-bottom: 0;" {
						td.DOM.Find("span").Remove()
						if i < len(headers) {
							header := headers[i]
							row[header] = strings.TrimSpace(strings.TrimSpace(td.DOM.Text()))
							//fmt.Println(row[header])
						}
					}
					// if td.Attr("class") == "mod-right" {
					// 	//fmt.Println("td.Text = ", strings.TrimSpace(tds))
					// 	if i < len(headers) {
					// 		header := headers[i]
					// 		row[header] = strings.TrimSpace(td.Text)
					// 	}
					//fmt.Println(td.DOM.Text())
					// tds := tr.DOM.Find("td")
					// tds.Find("span").Remove()
					// //fmt.Println(tds.Text())
					// if i < len(headers) {
					// 	header := headers[i]
					// 	row[header] = strings.TrimSpace(tds.Text())
					// }
					//}
					if td.Attr("class") == "mod-color-sq" {
						//fmt.Println(td.ChildText("span"))
						//fmt.Println(i, " ", len(headers))
						//if i < len(headers) {
						//fmt.Println(i)
						header := headers[i-3]
						//fmt.Println(headers[i-3])
						//fmt.Println(headers[i-3])
						row[header] = strings.TrimSpace(td.ChildText("span"))
						//fmt.Println(strings.TrimSpace(td.ChildText("span")))
						// if i < len(heajders) {
						// 	header := headers[i]
						// 	row[header] = strings.TrimSpace(td.Text)
						// }
						//}
					}
					//fmt.Println(table)
				})
				if len(row) > 0 {
					table = append(table, row)
					//fmt.Println(table)
				}
			}
		})
		//fmt.Println(table)
		// if len(table) > 0 {
		// 	tables = append(tables, table)
		// }
		for i := range table {
			for key, value := range table[i] {
				newValue := strings.ReplaceAll(value, "%", "")
				table[i][key] = newValue
			}
		}
		var saveMap map[string]string
		var testValMap string
		fmt.Scanln(&testValMap)
		copyNext := false
		mapWithIntValue := make(map[string]int)
		for i := range table {
			for key, value := range table[i] {
				if key == "Map" && value == testValMap {
					copyNext = true
				}
				if copyNext {
					delete(table[i], "Map")
					//fmt.Println(table[i])
					saveMap = table[i]
					copyNext = false
				}
			}

		}
		//fmt.Println(saveMap)
		for key, value := range saveMap {
			mapWithIntValue[key], _ = strconv.Atoi(value)
		}
		//fmt.Println(mapWithIntValue)
		result := sortMapByValues(mapWithIntValue)
		//fmt.Println(result)
		// for i := range result {
		// 	fmt.Println(i)
		// }
		for i := 0; i < 5; i++ {
			fmt.Print(result[i].Key, " ")
		}
	})

	fmt.Println("Scrapping Completed")

	c.Visit("https://www.vlr.gg/event/agents/2097/valorant-champions-2024")
	// fmt.Println(data)
}

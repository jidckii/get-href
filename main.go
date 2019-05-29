package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	scrapURL := flag.String("url", "https://github.com/PuerkitoBio/goquery", "URL for scraping")
	areaFilter := flag.String("area-filter", "head", "You filter for search href.")
	findFilter := flag.String("find-filter", "link[hreflang]", "You filter for search href.")
	outJSON := flag.Bool("json", false, "Response in json format.")
	flag.Parse()

	officeList := offoceFind(*scrapURL, *areaFilter, *findFilter)
	if len(officeList) == 0 {
		log.Println("Not found")
	} else {

		officeListJSON, err := json.Marshal(officeList)
		if err != nil {
			log.Fatal(err)
		}
		if *outJSON {
			fmt.Print(string(officeListJSON))
		} else {
			for _, v := range officeList {
				fmt.Print(v + " ")
			}
		}
	}
}

// Find all office path URLs
func offoceFind(scrapURL string, areaFilter string, findFilter string) (uris []string) {

	res, err := http.Get(scrapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find the review items
	doc.Find(areaFilter).Find(findFilter).Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			parseURL, err := url.Parse(href)
			if err != nil {
				log.Fatal(err)
			}
			pathURI := parseURL.EscapedPath()
			uris = append(uris, pathURI)
		}
	})
	return uris
}

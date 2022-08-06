package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/html"
)

type newResponse struct {
	Description string
	createdAt   interface{}
	Formated    interface{}
}

func main() {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9", nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println(response.Body)

	// data := parse(response.Body)

	newData := ParseNewFunc(response.Body)

	// for _, val := range newData {
	// 	fmt.Println(val)
	// }

	words := strings.SplitN(newData[5], "\n", -1)

	newWords := []string{}
	for _, val := range words {
		if len(val) > 3 {
			newWords = append(newWords, val)
		}
	}

	newResponses := []newResponse{}

	for i, val := range newWords {
		tmp := newResponse{}
		if i%2 != 0 {
			tmp.Description = val
			tmp.createdAt = newWords[i-1]
			tmp.Formated = map[string]interface{}{
				"createdAt": newWords[i-1],
			}
			newResponses = append(newResponses, tmp)
		}
	}
	// fmt.Println(newData[5])

	// return data,nil

}

func parse(text io.Reader) (data []string) {

	tkn := html.NewTokenizer(text)

	var vals []string

	var isLi bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			isLi = t.Data == "td"

		case tt == html.TextToken:

			t := tkn.Token()

			if isLi {
				vals = append(vals, t.Data)
			}

			isLi = false
		}
	}
}

func ParseNewFunc(text io.Reader) []string {
	doc, err := goquery.NewDocumentFromReader(text)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	newRow := []string{}
	doc.Find(".tracking").Children().Each(func(i int, sel *goquery.Selection) {
		rows := sel.Find("tr").Text()
		newRow = append(newRow, rows)
	})

	return newRow

}

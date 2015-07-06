package utils

import (
	"fmt"
	"meshalka/lib/config"
	"net/http"
	"regexp"
	"strings"
)

func AddHandler(page config.Page, fn config.Handler) {
	http.HandleFunc(page.Path, func(w http.ResponseWriter, r *http.Request) {
		if page.Args.Use {
			fmt.Println(page.Path)
			m := getUrlMatches(page, r.URL.Path)
			if m == nil {
				http.NotFound(w, r)
				return
			}
			fn(w, r, getPageWithArgs(page, m))
		} else {
			fn(w, r, page.Copy())
		}
	})
}

func getPageWithArgs(page config.Page, matches []string) config.Page {
	newPage := page.Copy()
	i := 1
	for arg := range page.Args.Con {
		newPage.Args.Con[arg] = matches[i]
		i++
	}

	return newPage
}

func getUrlMatches(page config.Page, path string) []string {
	regex := regexp.MustCompile(getUrlRegexStr(page))
	return regex.FindStringSubmatch(path)
}

func getUrlRegexStr(page config.Page) string {
	result := "^" + page.Path
	if len(page.Args.Con) > 0 {
		result = strings.TrimSuffix(result, "/")
	}
	for _, v := range page.Args.Con {
		result += "/(" + v + ")"
	}
	fmt.Println(result + "$")
	return result + "$"
}

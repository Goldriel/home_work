package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Dict struct {
	Word  string
	Count int
}

var re = regexp.MustCompile(`\s{2,}|\t{2,}|\n+`)

func Top10(str string) []string {
	// Place your code here.
	dict := make(map[string]int)
	var list []string
	var endList []string
	d := make([]Dict, 0)

	if str == "" {
		return nil
	}

	str = re.ReplaceAllString(str, " ")
	list = strings.Split(str, " ")

	for _, value := range list {
		_, ok := dict[value]
		if ok {
			dict[value]++
		} else {
			dict[value] = 1
		}
	}

	for word, count := range dict {
		d = append(d, Dict{word, count})
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i].Count > d[j].Count || (d[i].Count == d[j].Count && d[i].Word < d[j].Word)
	})

	for i := 0; ; i++ {
		if i == 10 || i == len(d) {
			break
		}
		endList = append(endList, d[i].Word)
	}
	return endList
}

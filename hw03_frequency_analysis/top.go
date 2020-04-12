package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

const WordsReturn = 10

// Top10 fucn retunr top10 words in s
func Top10(s string) []string {
	if s == "" {
		result := make([]string, 0)
		return result
	}
	type kv struct {
		Key   string
		Value int
	}
	ss := make([]kv, 100)
	result := make([]string, 0)
	counts := make(map[string]int)

	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")

	// reg, err := regexp.Compile("[^а-яА-Я0-9 ]+")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// s := reg.ReplaceAllString(s, "")

	for _, word := range strings.Fields(s) {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss[:WordsReturn] {
		result = append(result, kv.Key)
	}
	return result
}

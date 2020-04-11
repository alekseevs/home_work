package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"fmt"
	"strings"
	// "regexp"
	// "log"
	"sort"
)

// Top10 fucn retunr top10 words in s 
func Top10(s string) []string {
	
	counts := make(map[string]int)
	var result []string
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
    // reg, err := regexp.Compile("[^а-яА-Я0-9 ]+")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // s := reg.ReplaceAllString(s, "")

	for _,word := range strings.Fields(s) {
		_, ok := counts[word]
		if ok {
			counts[word] ++
		} else {
			counts[word]=1
		}
	}
	type kv struct {
        Key   string
        Value int
    }
	var ss []kv
	
	for k, v := range counts {
        ss = append(ss, kv{k, v})
    }
    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		result = append(result, kv.Key)
		if len(result) >= 10{
			break
		}
    }
	return result
}

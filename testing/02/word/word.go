package word

import (
	"log"
	"regexp"
	"strings"
)

func UsageCount(s string) map[string]int {
	xs := strings.Fields(removeNonAlphanumeric(s))
	m := make(map[string]int)
	for _, v := range xs {
		m[strings.ToLower(v)]++
	}
	return m
}

func TotalCount(s string) int {
	return len(strings.Fields(removeNonAlphanumeric(s)))
}

func removeNonAlphanumeric(s string) string {
	reg, err := regexp.Compile("\\W+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, " ")
}

package models

import "strings"

func HandLeTagsListData(tags []string) map[string]int {
	tagsMap := make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}

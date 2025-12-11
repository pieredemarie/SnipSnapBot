package services

import (
	"net/url"
	"strings"
)

func (s *LinkService) isValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

func (s *LinkService) parseTags(tagsStr string) []string {
	raw := strings.Fields(tagsStr)
	tagsMap := make(map[string]struct{})
	for _, t := range raw {
		t = strings.ToLower(t)
		tagsMap[t] = struct{}{}
	}
	tags := make([]string, 0, len(tagsMap))
	for t := range tagsMap {
		tags = append(tags, t)
	}
	return tags
}

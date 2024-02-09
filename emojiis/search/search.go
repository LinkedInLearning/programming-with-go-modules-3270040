package search

import (
	"slices"
	"strings"
)

// Params used to specify search parameters
type Params struct {
	Include []string // Include slice of strings to include in search
	Exclude []string // Exclude slice of strings to exclude in search
}

// ByDescription searches emojis using the params
func ByDescription(params Params) (result []Emoji) {

	for _, emo := range emojis {
		if shouldExclude(emo, params.Exclude) {
			continue
		}

		for _, include := range params.Include {
			include = strings.ToLower(include)
			if strings.Contains(emo.Label, include) || slices.Contains(emo.Tags, include) {
				result = append(result, emo)
			}
		}
	}

	return
}

// ByTags search for emojis with specified tags
func ByTags(tags ...string) (result []Emoji) {
	for _, emo := range emojis {
		for _, tag := range tags {
			if slices.Contains(emo.Tags, tag) {
				result = append(result, emo)
			}
		}
	}

	return
}

// shouldExclude checks emoji tags and labels for exclusions
func shouldExclude(emo Emoji, excludes []string) bool {
	for _, exclude := range excludes {
		exclude = strings.ToLower(exclude)
		if strings.Contains(emo.Label, exclude) || slices.Contains(emo.Tags, exclude) {
			return true
		}
	}

	return false
}

package search

import (
	"slices"
	"testing"
)

func TestByDesc(t *testing.T) {
	tests := map[string]struct {
		params Params
		want   []string
	}{
		"fruit": {
			Params{Include: []string{"fruit"}},
			[]string{"🍇", "🍈", "🍉", "🍊", "🍋"},
		},
		"cat": {
			Params{Include: []string{"cat"}},
			[]string{"🐱"},
		},
		"animal faces": {
			Params{Include: []string{"face"}, Exclude: []string{"smile", "laugh", "grin", "upside-down"}},
			[]string{"🐵", "🐶", "🐱", "🐯", "🦊"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ByDescription(test.params)
			if len(result) != len(test.want) {
				t.Errorf("Search ByDescription: want %d emojis, got %d", len(test.want), len(result))
			}

			for _, emoji := range result {
				if !slices.Contains(test.want, emoji.Emoji) {
					t.Errorf("Search ByDescription: expected %s to be in %#v", emoji, test.want)
				}
			}
		})
	}
}
